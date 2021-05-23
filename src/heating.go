package main

import (
	"github.com/labstack/gommon/log"
)

func heatingTask() {
	heatingConfig, err := LoadConfigheating()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	if heatingConfig.Status {
		suitUpArms(heatingConfig.Arms...)
		suitUpPearlArms(heatingConfig.PearlArms...)
		checkMainScreen()
		searchRepo()
		runHeating()
	}
}

func runHeating() {

}
