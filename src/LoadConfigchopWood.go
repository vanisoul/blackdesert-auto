package main

import (
	"github.com/spf13/viper"
)

func LoadConfigchopWood() (config ConfigDo, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("chopWood")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
