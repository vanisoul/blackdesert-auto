package main

import (
	"github.com/labstack/gommon/log"
)

func chopWoodTask() {
	heatingConfig, err := LoadConfigchopWood()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	processingTask(heatingConfig.Status, "chopWood", heatingConfig.Arms, heatingConfig.PearlArms, heatingConfig.Method)
}
