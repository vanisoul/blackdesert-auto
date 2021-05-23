package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func processingTask(status bool, typeStr string, arms []string, pearlArms []string, method []method) {

	if status {
		suitUpArms(arms...)
		checkMainScreen()
		suitUpPearlArms(pearlArms...)
		checkMainScreen()
		runTask(typeStr, method)
	}
}

func runTask(typeStr string, method []method) {
	// heatingConfig.Method[0].Formula[0].Name
	// heatingConfig.Method[0].Formula[0].Lov
	for _, med := range method {
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
			for proing {
				proing, _, _ = whilescreen("img/Processeding_1.png")
			}
			checkMainScreen()
			robotgo.KeyTap("space")
			robotgo.Sleep(3)
			searchRepo()
			saveRepoAll(insertStrToFilenameTailArr(med.Recycle, "bag")...)
		}
	}
}

func processPutAll(imgs ...string) {
	takeArticleSum := 0

	count := 3
	for count > 0 {
		for _, img := range imgs {
			succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "Formula"), 3)
			if succright {
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
