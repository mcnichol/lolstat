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

// Read this from config
var apiKey = getConfig()
var summoners = map[string]string{
	"dad":    "Buckethead Wendy",
	"august": "BatteryStaple123",
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

		resp, err := http.Get(fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summoners["dad"], apiKey))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		decoded := json.NewDecoder(bytes.NewReader(body))

		var s1 model.Summoner
		for {
			if err := decoded.Decode(&s1); err == io.EOF {
				break
			} else {
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		response := getRequest("https://na1.api.riotgames.com/lol/match/v4/timelines/by-match/3347123870?api_key=RGAPI-f06f4a32-71e4-4735-8370-7d8614027f80")

		devlog(string(response))
		devlog(string(body))
		p1.Text = s1.ToString()

		resp, err = http.Get(fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summoners["august"], apiKey))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		decoded = json.NewDecoder(bytes.NewReader(body))

		var s2 model.Summoner
		for {
			if err := decoded.Decode(&s2); err == io.EOF {
				break
			} else {
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		devlog(string(body))
		p2.Text = s2.ToString()
	}

	draw := func() {
		ui.Render(p1)
		ui.Render(p2)
	}

	draw()
	uiEvents := ui.PollEvents()

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

		default:
			updateParagraph()
			draw()
		}
	}

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
