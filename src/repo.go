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
	log.Info("searchRepo")
	robotgo.KeyTap("r")
	succCheckRepo, repox, repoy := whilescreen("img/bank.png", 2)
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
			robotgo.Sleep(60)
			robotgo.KeyTap("t")
			robotgo.Sleep(1)
			robotgo.KeyTap("control")
			robotgo.Sleep(1)
			robotgo.KeyTap("r")
			leftMosueforimg("img/bank.png")
		}
	}
}

func takeRepoOne(quantity int, imgs ...string) (succ bool) {
	oksucc, _, _ := whilescreen("img/bank_ok.png")

	if oksucc {
		count := 3
		for count > 0 {
			succarticle, articlex, articley := whilescreenManyEasy(3, insertStrToFilenameTailArr(imgs, "bank")...)
			if succarticle {
				rightMosue(articlex, articley)
				takeCount(quantity)
				leftMosueforimg("img/bagToBankEnter.png")
				robotgo.Sleep(1)
				return
			}
			scrollBankDown(8)
			count = count - 1
		}
	}
	return
}

func takeRepoAll(quantity int, imgs ...string) (succ bool) {
	oksucc, _, _ := whilescreen("img/bank_ok.png")
	takeArticleSum := 0
	if oksucc {
		count := 3
		for count > 0 {
			for _, img := range imgs {
				succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "bank"), 3)
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
			scrollBankDown(8)
			count = count - 1
		}
	}
	return
}

func saveRepoAll(imgs ...string) {
	oksucc, _, _ := whilescreen("img/bank_ok.png")
	takeArticleSum := 0
	if oksucc {
		count := 3
		for count > 0 {
			for _, img := range imgs {
				succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "bag"), 3)
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
			scrollBagDown(8)
			count = count - 1
		}
	}
}

func saveRepoOne(imgs ...string) {
	oksucc, _, _ := whilescreen("img/bank_ok.png")
	if oksucc {
		count := 3
		for count > 0 {
			succarticle, articlex, articley := whilescreenManyEasy(3, insertStrToFilenameTailArr(imgs, "bag")...)
			if succarticle {
				rightMosue(articlex, articley)
				leftMosueforimg("img/bag_max.png")
				leftMosueforimg("img/bagToBankEnter.png")
				robotgo.Sleep(1)
				return
			}
			scrollBagDown(8)
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
