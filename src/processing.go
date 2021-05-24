package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func processingTask(status bool, typeStr string, arms []string, pearlArms []string, method []method) {

	if status {
		infoConfig, err := LoadConfigInfo()
		if err != nil {
			log.Errorf("cannot load config:", err)
			return
		}
		if infoConfig.ProcessingCutoverArm {
			suitUpArms(arms...)
			checkMainScreen()
			suitUpPearlArms(pearlArms...)
			checkMainScreen()
		}
		runTask(typeStr, method)
	}
}

func runTask(typeStr string, method []method) {
	// heatingConfig.Method[0].Formula[0].Name
	// heatingConfig.Method[0].Formula[0].Lov
	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	tmpDrinkToWork := 0
	methodnumbers := generateRandomNumber(0, len(method), len(method))
	for _, mednum := range methodnumbers {
		med := method[mednum]
		if tmpDrinkToWork == infoConfig.DrinkingOnTheWayToWork {
			checkMainScreen()
			beerTask()
			tmpDrinkToWork = 0
		}
		succCount := checkCount(med.Formula)
		if succCount {

			searchRepo()
			leftMosueforimg("img/ProcessingButton.png")
			isStr := fmt.Sprintf("img/is%s.png", typeStr)
			notStr := fmt.Sprintf("img/not%s.png", typeStr)
			gui, _, _ := whilescreenEasy(isStr, 3)
			if !gui {
				leftMosueforimgEasy(notStr)
			}
			for _, fml := range med.Formula {
				processPutAll(fml.Name)
			}
			robotgo.Sleep(1)
			leftMosueforimg("img/ProcessingStart.png")
			proing := true

			tmpTimeSec := 0
			for proing {
				proing = screenYesOrNoEasy("img/Processeding_1.png", 20)
				if proing {
					tmpTimeSec = tmpTimeSec + 1
				} else {
					tmpTimeSec = tmpTimeSec + 20
					checkMainScreen()
					tmpTimeSec = tmpTimeSec + 3
					proing = screenYesOrNoEasy("img/Processeding_1.png", 20)
					tmpTimeSec = tmpTimeSec + 1
				}
				if tmpTimeSec == infoConfig.ProcessedTimeSec {
					break
				}
			}
			checkMainScreen()
			robotgo.KeyTap("space")
			robotgo.Sleep(3)
			searchRepo()
			saveRepoAll(med.Recycle...)
			tmpDrinkToWork = tmpDrinkToWork + 1
		}
	}
}

func processPutAll(imgs ...string) {
	takeArticleSum := 0
	count := 3
	for count > 0 {
		for _, img := range imgs {
			succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "Formula"), 6)
			if succright {
				robotgo.Sleep(1)
				takeArticleSum = takeArticleSum + 1
			}
			if takeArticleSum == len(imgs) {
				log.Info("takeArticleSum :", takeArticleSum)
				log.Info("len(imgs) :", len(imgs))
				return
			}
		}
		scrollProcessDown(8)
		count = count - 1
	}

}

func scrollProcessDown(fre int) {
	_, x, y := whilescreen("img/ProcessingOK.png")
	log.Infof("scrollProcessDown x:%d, y:%d", x, y)
	robotgo.MoveMouse(x+41, y+103)
	scrolldown(fre)
}
