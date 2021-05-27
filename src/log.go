package main

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func setLog(typeStr string, msg string, item string) {
	CurrentTime := time.Now().Format("2006-01-02 15:04:05")
	vipStatusLog := viper.New()
	vipStatusLog.AddConfigPath("./log/")
	vipStatusLog.SetConfigName("log")
	vipStatusLog.SetConfigType("json")

	originData, err := LoadConfiglog()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	newDatas := originData.Log

	var data Log
	data.CurrentTime = CurrentTime
	data.Item = item
	data.Msg = msg
	data.TypeStr = typeStr
	newDatas = append(newDatas, data)
	vipStatusLog.AutomaticEnv()
	vipStatusLog.Set("log", newDatas)

	vipStatusLog.WriteConfig()
}

func LoadConfiglog() (config ConfigLog, err error) {
	exis := FileExist("log/log.json")
	if !exis {
		err := CopyFile("log/log-defult.json", "log/log.json")
		if err != nil {
			fmt.Printf("CopyFile failed %q\n", err)
		} else {
			fmt.Printf("CopyFile succeeded\n")
		}
	}

	viper.AddConfigPath("./log/")
	viper.SetConfigName("log")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ConfigLog struct {
	Log []Log `mapstructure:"log"`
}

type Log struct {
	CurrentTime string `mapstructure:"currentTime"`
	TypeStr     string `mapstructure:"typeStr"`
	Msg         string `mapstructure:"msg"`
	Item        string `mapstructure:"item"`
}
