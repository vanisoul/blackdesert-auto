package main

import (
	"github.com/spf13/viper"
)

func LoadConfigAccount() (config Config, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("account")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type Config struct {
	Account  string `mapstructure:"account"`
	Password string `mapstructure:"password"`
}
