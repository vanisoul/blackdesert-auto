package main

import (
	"github.com/spf13/viper"
)

func LoadConfigMoveItems() (config ConfigMoveItems, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("move-items")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type ConfigMoveItems struct {
	Status bool        `mapstructure:"status"`
	Moves  []MoveItems `mapstructure:"moves"`
}

type MoveItems struct {
	HaulType       string   `mapstructure:"haulType"`
	Source         string   `mapstructure:"source"`
	SourceSeq      int      `mapstructure:"sourceSeq"`
	Destination    string   `mapstructure:"destination"`
	DestinationSeq int      `mapstructure:"destinationSeq"`
	Items          []string `mapstructure:"items"`
}
