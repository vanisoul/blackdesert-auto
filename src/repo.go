package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func checkCount(fmls []formula) (succ bool) {
	searchRepo()
	oksucc, _, _ := whilescreen("img/bank_ok.png")
	takeArticleSum := 0
	if oksucc {
		count := 3
		for count > 0 {
			for _, fml := range fmls {
				succright := rightMosueforimgEasy(insertStrToFilenameTail(fml.Name, "bank"), 3)
				if succright {
					takeCount(fml.Lov)
					LoVStr := fmt.Sprintf("img/%d_VoL.png", fml.Lov)
					LoVsucc, _, _ := whilescreenEasy(LoVStr, 3)
					if LoVsucc {
						setLog("checkCount", "加工材料數量合格", fml.Name)
						robotgo.KeyTap("esc")
					} else {
						setLog("checkCount", "加工材料數量不合格 離開此配方加工", fml.Name)
						robotgo.KeyTap("esc")
						robotgo.Sleep(1)
						succ = false
						return
					}
					robotgo.Sleep(1)
					takeArticleSum = takeArticleSum + 1
				}
				if takeArticleSum == len(fmls) {
					succ = true
					return
				}
			}
			scrollBankDown(8)
			count = count - 1
		}
	}
	return
}

func searchRepo() {
	checkMainScreen()
	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	log.Info("searchRepo")
	robotgo.KeyTap("r")
	succCheckRepo, repox, repoy := whilescreen("img/bank.png", 2)
	if succCheckRepo {
		setLog("searchRepo", "原地尋找到倉庫", "")
		leftMosue(repox, repoy)
	} else {
		checkMainScreen()
		setLog("searchRepo", "需要回家再走過去", "")
		// TODO : 先回家 在走到NPC 比較不會有誤差
		// goToHome()
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
			robotgo.KeyTap("t")
			robotgo.Sleep(1)
			robotgo.KeyTap("control")
			robotgo.Sleep(1)
			robotgo.KeyTap("r")
			leftBanksucc := leftMosueforimg("img/bank.png")
			if leftBanksucc {
				setLog("searchRepo", "成功開啟倉庫", "")
			}
		}
	}
}

func takeRepoOne(quantity int, imgs ...string) (succ bool) {
	oksucc, _, _ := whilescreen("img/bank_ok.png")
	succ = false
	if oksucc {
		count := 3
		for count > 0 {
			newImgs := insertStrToFilenameTailArr(imgs, "bank")
			succarticle, articlex, articley := whilescreenManyEasy(3, newImgs...)
			if succarticle {
				rightMosue(articlex, articley)
				takeCount(quantity)
				leftMosueforimg("img/bagToBankEnter.png")
				robotgo.Sleep(1)
				succ = true
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
					setLog("saveRepoAll", "成功存放物品", img)

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

func saveRepoOne(imgs ...string) (succ bool) {
	succ = false
	oksucc, _, _ := whilescreen("img/bank_ok.png")
	if oksucc {
		count := 3
		for count > 0 {
			succarticle, articlex, articley := whilescreenManyEasy(3, insertStrToFilenameTailArr(imgs, "bag")...)
			if succarticle {
				rightMosue(articlex, articley)
				leftMosueforimg("img/bag_max.png")
				leftMosueforimg("img/bagToBankEnter.png")
				setLog("saveRepoOne", "成功存放其一物品", strings.Join(imgs, ", "))

				robotgo.Sleep(1)
				succ = true
				return
			}
			scrollBagDown(8)
			count = count - 1
		}
	}
	return
}

func takeCount(quantity int) {
	robotgo.Sleep(1)
	for _, key := range strconv.Itoa(quantity) {
		robotgo.KeyTap(string(key))
		robotgo.Sleep(1)
	}
}
