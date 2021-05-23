package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func suitUpArms(arms ...string) {
	heatingConfig, err := LoadConfigheating()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	robotgo.Sleep(1)
	succarmsui, _, _ := whilescreen("img/arms_ok.png")
	if succarmsui {
		count := 3
		for count > 0 {
			robotgo.Sleep(1)
			rightMosueforimgEasyAll(heatingConfig.Arms, 3)
			scrollBagDown(8)
			count = count - 1
		}
	} else {
		robotgo.KeyTap("i")
		suitUpArms(arms...)
	}

}

func suitUpPearlArms(pearlArms ...string) {

}
