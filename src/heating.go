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
	setLog("heatingTask", "開始加熱", "Heating")
	processingTask(heatingConfig.Status, "Heating", heatingConfig.Arms, heatingConfig.PearlArms, heatingConfig.Method)
}
