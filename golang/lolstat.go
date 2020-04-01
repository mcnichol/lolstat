package main

import (
	"./model"
	"bytes"
	"encoding/json"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	BaseUrl = "https://na1.api.riotgames.com/lol"
)

var apiKey = getConfig()
var summoners = map[string]string{
	"dad":    "Buckethead Wendy",
	"august": "BatteryStaple123",
}

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var p1 = widgets.NewParagraph()
	var p2 = widgets.NewParagraph()

	updateParagraph := func() {
		var cmd = exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}

		strOut := string(out)
		strArray := strings.Fields(strOut)
		strX := strArray[1]
		strY := strArray[0]

		rectX, err := strconv.ParseInt(strX, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		rectY, err := strconv.ParseInt(strY, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		p1.SetRect(0, 0, int(rectX), int(rectY)/2)
		p2.SetRect(0, int(rectY)/2, int(rectX), int(rectY))
	}

	refreshData := func() {
		dad := getSummoner(summoners["dad"])
		devlog(dad.ToString())

		august := getSummoner(summoners["august"])
		devlog(august.ToString())

		p1.Text = fmt.Sprintf("Summoner Info:\n\n%15v:%50s\n%15v:%50d\n%15v:%50v\n%15v:%50v\n",
			"Name", dad.Name, "Level", dad.Level, "Account ID", dad.AccountId, "Summoner ID", dad.Id)

		p2.Text = fmt.Sprintf("Summoner Info:\n\n%15v:%50s\n%15v:%50d\n%15v:%50v\n%15v:%50v\n",
			"Name", august.Name, "Level", august.Level, "Account ID", august.AccountId, "Summoner ID", august.Id)

		//respMatchTimeline := getRequest(fmt.Sprintf(BASE_URL + "/match/v4/timelines/by-match/%s?api_key=%s", "3347123870", apiKey))
	}

	draw := func() {
		ui.Render(p1)
		ui.Render(p2)
	}

	draw()
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	ticker2 := time.NewTicker(30 * time.Second).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "r":
				refreshData()
				draw()
			}
		case <-ticker:
			updateParagraph()
			draw()
		case <-ticker2:
			ml := getMatchList("K8tn-mm47NcDCA0-oMh7buzFhkLhPQ1lUWcjGLYoqi1pEw")
			devlog(ml.ToStringSlice(2))
		}
	}
}

func getConfig() string {
	devlog("Reading API key from file")
	file, _ := ioutil.ReadFile("../config/riot-api.key")
	if string(file[0:5]) != "RGAPI" {
		devlog("This appears to be an invalid Riot API Key Structure, check for proper key in `root_dir/config/riot-api.key`")
		log.Fatal("Incorrect Riot API Key")
	}
	return string(file)
}

func devlog(data string) () {
	f, err := os.OpenFile("local.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(time.Now().String() + "\t" + data + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func getSummoner(summonerName string) *model.Summoner {
	var s model.Summoner

	respSummoner := getRequest(fmt.Sprintf(BaseUrl+"/summoner/v4/summoners/by-name/%s?api_key=%s", summonerName, apiKey))
	decoded := json.NewDecoder(bytes.NewReader(respSummoner))

	for {
		if err := decoded.Decode(&s); err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return &s
}

func getMatchList(accountId string) *model.MatchList {
	var ml model.MatchList

	respMatchList := getRequest(fmt.Sprintf(BaseUrl+"/match/v4/matchlists/by-account/%s?api_key=%s", accountId, apiKey))
	decoded := json.NewDecoder(bytes.NewReader(respMatchList))

	for {
		if err := decoded.Decode(&ml); err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return &ml
}

func getMatch(matchId uint64) *model.Match {
	var m model.Match

	respSummoner := getRequest(fmt.Sprintf(BaseUrl+"/match/v4/matches/by-name/%d?api_key=%s", matchId, apiKey))
	decoded := json.NewDecoder(bytes.NewReader(respSummoner))

	for {
		if err := decoded.Decode(&m); err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return &m
}

func getRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
