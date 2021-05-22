package main

import (
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func searchRepo() {
	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	key := strings.Split(infoConfig.SearchNPC, "+")[1]
	robotgo.KeyTap(key, "alt")
	sersucc, x, y := whilescreen("img/searchNPC.png")
	if sersucc {
		textLocation(x-240, y, infoConfig.Repo)
		leftMosue(x, y)
		leftMosue(x-240, y-170)
		robotgo.Sleep(1)
		robotgo.KeyTap("t")
		robotgo.Sleep(20)
		robotgo.KeyTap("t")
		robotgo.Sleep(1)
		robotgo.KeyTap("r")
		leftMosueforimg("img/bank.png")
	}
}

func tackRepo(img string) {

}

func saveRepo(img string) {

}
