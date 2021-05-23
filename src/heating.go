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
	processingTask(heatingConfig.Status, "Heating", heatingConfig.Arms, heatingConfig.PearlArms, heatingConfig.Method)
}
