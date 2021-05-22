package main

import (
	"github.com/labstack/gommon/log"
)

func main() {
	openGameClient()
	accConfig, err := LoadConfigAccount()
	if err != nil {
		log.Errorf("cannot load config:", err)
	}
	log.Infof(accConfig.Account)
	log.Infof(accConfig.Password)
}
