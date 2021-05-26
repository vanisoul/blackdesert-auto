package main

import (
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
	viper.AddConfigPath("./status/")
	viper.SetConfigName("status")
	viper.SetConfigType("json")

	schedule := []int{}
	viper.AutomaticEnv()
	viper.Set("status", true)
	viper.Set("type", typeName)
	viper.Set("schedule", schedule)
	viper.Set("now", 0)
	viper.WriteConfig()
}

func setStatusMethod(typeName string, methodnumbers []int, count int) {
	viper.AddConfigPath("./status/")
	viper.SetConfigName("status")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	viper.Set("status", true)
	viper.Set("type", typeName)
	viper.Set("schedule", methodnumbers)
	viper.Set("now", count)
	viper.WriteConfig()
}

func endStatus() {
	viper.AddConfigPath("./status/")
	viper.SetConfigName("status")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	viper.Set("status", false)
	viper.Set("type", "")
	viper.Set("schedule", []int{})
	viper.Set("now", -1)
	viper.WriteConfig()
}
