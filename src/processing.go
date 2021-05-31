package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func processingTask(status bool, typeStr string, arms []string, pearlArms []string, method []method, mattersArr ...bool) {

	if status {
		matters := true
		if len(mattersArr) != 0 {
			matters = mattersArr[0]
		}

		infoConfig, err := LoadConfigInfo()
		if err != nil {
			log.Errorf("cannot load config:", err)
			return
		}

		if infoConfig.ProcessingCutoverArm && checkStatusArms(typeStr) {
			setLog("processingTask", "開始換一般裝備", strings.Join(arms, ", "))
			suitUpArms(arms...)
			checkMainScreen()
			setLog("processingTask", "開始換珍珠裝備", strings.Join(pearlArms, ", "))
			suitUpPearlArms(pearlArms...)
			checkMainScreen()
			setStatusArms(typeStr)
		}

		runTask(typeStr, method, matters)
	}
}

func runTask(typeStr string, method []method, matters bool) {

	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	tmpDrinkToWork := 0
	succ, methodnumbers, skip := checkStatusMethodnumbersCount(typeStr, method)
	if !succ {
		log.Errorf("error checkStatusMethodnumbersCount", err)
		return
	}
	setLog("runTask", "配方優先順序", arrayToString(methodnumbers, ", "))
	setLog("runTask", "配方跳過數量", strconv.Itoa(skip))
	tmpcount := 0
	for _, mednum := range methodnumbers {
		if tmpcount < skip {
			setLog("runTask", "配方跳過 目前序", strconv.Itoa(tmpcount))
			tmpcount = tmpcount + 1
			continue
		}

		setStatusMethod(typeStr, methodnumbers, tmpcount)
		med := method[mednum]
		if tmpDrinkToWork == infoConfig.DrinkingOnTheWayToWork {
			checkMainScreen()
			beerTask()
			tmpDrinkToWork = 0
		}
		setLog("runTask", "本輪配方", FormulaNameArrayToString(med.Formula, " + "))
		succCount := checkCount(med.Formula)
		if succCount {
			searchRepo()
			leftMosueforimg("img/ProcessingButton.png")
			isStr := fmt.Sprintf("img/is%s.png", typeStr)
			notStr := fmt.Sprintf("img/not%s.png", typeStr)
			setLog("runTask", "確認加工類別", typeStr)
			gui, _, _ := whilescreenEasy(isStr, 3)
			if !gui {
				setLog("runTask", "選擇該加工類別", typeStr)
				entypesucc := leftMosueforimgEasy(notStr)
				if !entypesucc {
					return
				}
			}
			fmlNames := []string{}
			for _, fml := range med.Formula {
				fmlNames = append(fmlNames, fml.Name)
			}
			takeArticleSum := processPutAll(fmlNames...)
			if takeArticleSum == len(med.Formula) {
				r1 := saveIMG(true)
				setLog("runTask", "開始加工", strings.Join(med.Recycle, " ,"))
				setLog("runTask", "圖片", strconv.Itoa(r1))
				stSucc, stx, sty := whilescreenMany(20, "img/ProcessingStart.png", "img/ProcessingStartOne.png")
				if stSucc {
					r2 := saveIMG(true)
					setLog("runTask", "圖片", strconv.Itoa(r2))
					leftMosue(stx, sty)
				}
			}

			robotgo.Sleep(1)

			proing := true

			tmpTimeSec := 0
			setLog("runTask", "加工開始", strconv.Itoa(tmpTimeSec))
			for proing {
				proing = screenYesOrNoEasy("img/Processeding_1.png", 20)

				if proing {
					if tmpTimeSec == 0 {
						r := saveIMG(true)
						setLog("runTask", "開始加工", strconv.Itoa(r))
						if matters {
							setLog("additionalMatters", "計時等於0時執行並行額外工作", "")
							additionalMatters()
						}
					}
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
			setLog("runTask", "加工結束", strconv.Itoa(tmpTimeSec))
			checkMainScreen()
			robotgo.KeyTap("space")
			robotgo.Sleep(3)
			searchRepo()
			setLog("runTask", "成本放入倉庫", strings.Join(med.Recycle, " ,"))
			saveRepoAll(med.Recycle...)
			tmpDrinkToWork = tmpDrinkToWork + 1
		}
		tmpcount = tmpcount + 1
	}
	setStatusMethod(typeStr, methodnumbers, -2)
}

func processPutAll(imgs ...string) (takeArticleSum int) {
	takeArticleSum = 0
	count := 3
	for count > 0 {
		for _, img := range imgs {
			succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "Formula"), 6)
			if succright {
				robotgo.Sleep(1)
				takeArticleSum = takeArticleSum + 1
				setLog("processPutAll", "成功放入加工材料", img)
				setLog("processPutAll", "目前放入數量", strconv.Itoa(takeArticleSum))
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
	return
}

func scrollProcessDown(fre int) {
	_, x, y := whilescreen("img/ProcessingOK.png")
	log.Infof("scrollProcessDown x:%d, y:%d", x, y)
	robotgo.MoveMouse(x+41, y+103)
	scrolldown(fre)
}
