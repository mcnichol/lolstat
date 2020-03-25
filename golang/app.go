package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
// Read this from config
const apiKey = ""

func main() {
	summoners :=	map[string] string{
		"dad" : "Buckethead Wendy",
		"august" : "BatteryStaple123",
	}

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
		p1.Text = string(body)

		resp, err = http.Get(fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summoners["august"], apiKey))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		p2.Text = string(body)
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
