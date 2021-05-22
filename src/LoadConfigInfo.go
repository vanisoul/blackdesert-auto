package main

import (
	"github.com/spf13/viper"
)

func LoadConfigInfo() (config ConfigInfo, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("info")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ConfigInfo struct {
	Role     string `mapstructure:"role"`
	Location string `mapstructure:"location"`
	GamePath string `mapstructure:"gamePath"`
}
