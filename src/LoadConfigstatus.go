package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfigstatus() (config ConfigStatus, err error) {
	exis := FileExist("status/status.json")
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

type ConfigStatus struct {
	Status   bool   `mapstructure:"status"`
	Type     string `mapstructure:"type"`
	Schedule []int  `mapstructure:"schedule"`
	Now      int    `mapstructure:"now"`
}
