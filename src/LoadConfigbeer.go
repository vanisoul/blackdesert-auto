package main

import (
	"github.com/spf13/viper"
)

func LoadConfigbeer() (config ConfigBeer, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("beer")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ConfigBeer struct {
	Status       bool     `mapstructure:"status"`
	HotKey       string   `mapstructure:"hotKey"`
	ArticlesTake []string `mapstructure:"articlesTake"`
	ArticlesSave []string `mapstructure:"articlesSave"`
	Count        int      `mapstructure:"count"`
}
