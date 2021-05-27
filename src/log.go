package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func setLog(typeStr string, msg string, item string) {
	currentTime := time.Now()
}

func LoadConfiglog() (config ConfigStatus, err error) {
	exis := FileExist("log/log.json")
	if !exis {
		err := CopyFile("status/status-defult.json", "status/status.json")
		if err != nil {
			fmt.Printf("CopyFile failed %q\n", err)
		} else {
			fmt.Printf("CopyFile succeeded\n")
		}
	}

	viper.AddConfigPath("./status/")
	viper.SetConfigName("status")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
