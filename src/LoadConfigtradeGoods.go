package main

import (
	"github.com/spf13/viper"
)

func LoadConfigtradeGoods() (config ConfigtradeGoods, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("tradeGoods")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ConfigtradeGoods struct {
	Status  bool     `mapstructure:"status"`
	City    string   `mapstructure:"city"`
	Seq     int      `mapstructure:"seq"`
	Mineral []string `mapstructure:"mineral"`
	Wood    []string `mapstructure:"wood"`
}
