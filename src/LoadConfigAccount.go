package main

import (
	"github.com/spf13/viper"
)

func LoadConfigAccount() (config ConfigAccount, err error) {
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

type ConfigAccount struct {
	Account  string `mapstructure:"account"`
	Password string `mapstructure:"password"`
	FAkey    string `mapstructure:"2FAkey"`
}
