package main

import (
	"strconv"
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

func tackRepoOne(quantity int, imgs ...string) (succ bool) {
	oksucc, x, y := whilescreen("img/bank_ok.png")

	if oksucc {
		count := 3
		for count > 0 {
			succarticle, articlex, articley := whilescreenMany(imgs...)
			if succarticle {
				rightMosue(articlex, articley)
				takeCount(quantity)
				leftMosueforimg("img/bagToBankEnter.png")
				return
			}
			robotgo.MoveMouse(x-148, y+131)
			scrolldown(6)
			count = count - 1
		}
	}
	return
}

func tackRepoAll(quantity int, imgs ...string) (succ bool) {
	oksucc, x, y := whilescreen("img/bank_ok.png")

	if oksucc {
		count := 3
		for count > 0 {
			for _, img := range imgs {
				succright := rightMosueforimg(img, 10)
				if succright {
					takeCount(quantity)
					// leftMosueforimg("img/bag_max.png")
					leftMosueforimg("img/bagToBankEnter.png")
				}
			}
			robotgo.MoveMouse(x-148, y+131)
			scrolldown(6)
			count = count - 1
		}
	}
	return
}

func saveRepoAll(imgs ...string) {
	oksucc, x, y := whilescreen("img/bank_ok.png")

	if oksucc {
		count := 3
		for count > 0 {
			for _, img := range imgs {
				succright := rightMosueforimg(img, 10)
				if succright {
					leftMosueforimg("img/bag_max.png")
					leftMosueforimg("img/bagToBankEnter.png")
				}
			}
			robotgo.MoveMouse(x-3, y+151)
			scrolldown(8)
			count = count - 1
		}
	}
}

func takeCount(quantity int) {
	robotgo.Sleep(1)
	for _, key := range strconv.Itoa(quantity) {
		robotgo.KeyTap(string(key))
		robotgo.Sleep(1)
	}
}
