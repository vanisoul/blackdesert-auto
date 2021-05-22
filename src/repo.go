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
	checkMainScreen()
	robotgo.KeyTap("r")
	succCheckRepo, repox, repoy := whilescreen("img/bank.png")
	if succCheckRepo {
		leftMosue(repox, repoy)
	} else {
		checkMainScreen()
		key := strings.Split(infoConfig.SearchNPC, "+")[1]
		robotgo.KeyTap(key, "alt")
		sersucc, x, y := whilescreen("img/searchNPC.png")
		if sersucc {
			textLocation(x-240, y, infoConfig.Repo)
			leftMosue(x, y)
			leftMosue(x-240, y-170)
			robotgo.Sleep(1)
			robotgo.KeyTap("t")
			robotgo.Sleep(5)
			robotgo.KeyTap("t")
			robotgo.Sleep(1)
			robotgo.KeyTap("control")
			robotgo.Sleep(1)
			robotgo.KeyTap("r")
			leftMosueforimg("img/bank.png")
		}
	}
}

func tackRepoOne(quantity int, imgs ...string) (succ bool) {
	oksucc, x, y := whilescreen("img/bank_ok.png")

	if oksucc {
		count := 3
		for count > 0 {
			succarticle, articlex, articley := whilescreenManyEasy(3, imgs...)
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
	takeArticleSum := 0
	if oksucc {
		count := 3
		for count > 0 {
			for _, img := range imgs {
				succright := rightMosueforimgEasy(img, 3)
				if succright {
					takeCount(quantity)
					// leftMosueforimg("img/bag_max.png")
					leftMosueforimg("img/bagToBankEnter.png")
					robotgo.Sleep(1)
					takeArticleSum = takeArticleSum + 1
				}
				if takeArticleSum == len(imgs) {
					return
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
	takeArticleSum := 0
	if oksucc {
		count := 3
		for count > 0 {
			for _, img := range imgs {
				succright := rightMosueforimgEasy(img, 3)
				if succright {
					leftMosueforimg("img/bag_max.png")
					leftMosueforimg("img/bagToBankEnter.png")
					robotgo.Sleep(1)
					takeArticleSum = takeArticleSum + 1
				}
				if takeArticleSum == len(imgs) {
					return
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
