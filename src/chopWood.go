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
	setLog("chopWoodTask", "開始砍柴", "ChopWood")
	processingTask(heatingConfig.Status, "ChopWood", heatingConfig.Arms, heatingConfig.PearlArms, heatingConfig.Method)
}
