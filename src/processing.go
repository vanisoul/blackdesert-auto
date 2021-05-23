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
			isStr := fmt.Sprintf("img/is%s.img", typeStr)
			notStr := fmt.Sprintf("img/not%s.img", typeStr)
			gui, _, _ := whilescreen(isStr, 3)
			if !gui {
				leftMosueforimg(notStr)
			}
			for _, fml := range med.Formula {
				rightMosueforimgEasy(insertStrToFilenameTail(fml.Name, "Formula"), 5)
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
