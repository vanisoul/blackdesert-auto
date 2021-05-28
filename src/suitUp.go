package main

import (
	"github.com/go-vgo/robotgo"
)

func suitUpArms(arms ...string) {
	robotgo.Sleep(1)
	succarmsui, _, _ := whilescreenEasy("img/arms_ok.png", 3)
	if !succarmsui {
		robotgo.KeyTap("i")
	}
	succbagui, _, _ := whilescreenEasy("img/isPearl.png", 3)
	if succbagui {
		leftMosueforimgEasy("img/notBag.png")
	}
	count := 3
	tmpfre := 0
	for count > 0 {
		_, fre := rightMosueforimgEasyAll(setArmsPathArr(arms), 10)
		tmpfre = tmpfre + fre
		if tmpfre == len(arms) {
			return
		}
		scrollBagDown(8)
		count = count - 1
	}

}

func suitUpPearlArms(pearlArms ...string) {
	robotgo.Sleep(1)
	succarmsui, _, _ := whilescreenEasy("img/arms_ok.png", 3)
	if !succarmsui {
		robotgo.KeyTap("i")
	}
	succbagui, _, _ := whilescreenEasy("img/isBag.png", 3)
	if succbagui {
		leftMosueforimgEasy("img/notPearl.png")
	}
	count := 3
	tmpfre := 0
	for count > 0 {
		_, fre := rightMosueforimgEasyAll(setPearlArmsPathArr(pearlArms), 3)
		tmpfre = tmpfre + fre
		if tmpfre == len(pearlArms) {
			return
		}
		scrollBagDown(8)
		count = count - 1
	}

}
