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
	Role           int    `mapstructure:"role"`
	Location       string `mapstructure:"location"`
	GamePath       string `mapstructure:"gamePath"`
	SelectGraphics int    `mapstructure:"selectGraphics"`
	GameScreenX    int    `mapstructure:"gameScreenX"`
	GameScreenY    int    `mapstructure:"gameScreenY"`
	GameScreenW    int    `mapstructure:"gameScreenW"`
	GameScreenH    int    `mapstructure:"gameScreenH"`
	SearchNPC      string `mapstructure:"searchNPC"`
	Repo           string `mapstructure:"repo"`
}
