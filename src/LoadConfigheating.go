package main

import (
	"github.com/spf13/viper"
)

func LoadConfigheating() (config ConfigHeating, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("heating")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ConfigHeating struct {
	Status    bool     `mapstructure:"status"`
	Arms      []string `mapstructure:"arms"`
	PearlArms []string `mapstructure:"pearlArms"`
	Method    []method `mapstructure:"method"`
}

type formula struct {
	Name string `mapstructure:"name"`
	Lov  string `mapstructure:"LoV"`
}

type method struct {
	Formula []formula `mapstructure:"formula"`
	Recycle []string  `mapstructure:"recycle"`
}
