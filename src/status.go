package main

import (
	"strconv"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// func getToBeer() (succ bool, now int) {
// 	statusConfig, err := LoadConfigstatus()
// 	now = -1
// 	if err != nil {
// 		log.Errorf("cannot load status config:", err)
// 		succ = false
// 		return
// 	}
// 	if statusConfig.Status && statusConfig.Type == "beer" {
// 		succ = true
// 		now = statusConfig.Now
// 		return
// 	}
// 	return
// }

//傳入哪種加工
func getToProcessing(typeName string) (succ bool, now int) {
	statusConfig, err := LoadConfigstatus()
	now = -1
	if err != nil {
		log.Errorf("cannot load status config:", err)
		succ = false
		return
	}
	if statusConfig.Status && statusConfig.Type == typeName {
		succ = true
		now = statusConfig.Now
		return
	}
	return
}

func checkStatusArms(typeStr string) (succ bool) {
	statusConfig, err := LoadConfigstatus()
	if err != nil {
		log.Errorf("cannot load status config:", err)
		succ = false
		return
	}
	if statusConfig.Status && statusConfig.Type == typeStr && statusConfig.Now < 0 {
		succ = true
		return
	}
	return
}

func checkStatusMethodnumbersCount(typeStr string, method []method) (succ bool, methodnumbers []int, skip int) {
	statusConfig, err := LoadConfigstatus()
	if err != nil {
		log.Errorf("cannot load status config:", err)
		succ = false
		methodnumbers = generateRandomNumber(0, len(method), len(method))
		skip = 0
		return
	}
	if statusConfig.Status && statusConfig.Type == typeStr && statusConfig.Now > -1 && len(statusConfig.Schedule) != 0 {
		succ = true
		methodnumbers = statusConfig.Schedule
		skip = statusConfig.Now
		return
	}
	succ = true
	methodnumbers = generateRandomNumber(0, len(method), len(method))
	skip = 0
	return
}

/**
  "status":true,
  "type":"heating",
  "schedule":[0,1,4,2,3],
  "now":4
*/
func setStatusArms(typeName string) {
	vipStatusArms := viper.New()
	vipStatusArms.AddConfigPath("./status/")
	vipStatusArms.SetConfigName("status")
	vipStatusArms.SetConfigType("json")

	schedule := []int{}
	vipStatusArms.AutomaticEnv()
	vipStatusArms.Set("status", true)
	vipStatusArms.Set("type", typeName)
	vipStatusArms.Set("schedule", schedule)
	vipStatusArms.Set("now", 0)
	vipStatusArms.WriteConfig()
}

func setStatusMethod(typeName string, methodnumbers []int, count int) {
	setLog("setStatusMethod", "儲存執行位置", strconv.Itoa(count))
	vipStatusMethod := viper.New()
	vipStatusMethod.AddConfigPath("./status/")
	vipStatusMethod.SetConfigName("status")
	vipStatusMethod.SetConfigType("json")

	vipStatusMethod.AutomaticEnv()
	vipStatusMethod.Set("status", true)
	vipStatusMethod.Set("type", typeName)
	vipStatusMethod.Set("schedule", methodnumbers)
	vipStatusMethod.Set("now", count)
	vipStatusMethod.WriteConfig()
}

func endStatus() {
	vipStatusEnd := viper.New()
	vipStatusEnd.AddConfigPath("./status/")
	vipStatusEnd.SetConfigName("status")
	vipStatusEnd.SetConfigType("json")

	vipStatusEnd.AutomaticEnv()
	vipStatusEnd.Set("status", false)
	vipStatusEnd.Set("type", "")
	vipStatusEnd.Set("schedule", []int{})
	vipStatusEnd.Set("now", -1)
	vipStatusEnd.WriteConfig()
}
