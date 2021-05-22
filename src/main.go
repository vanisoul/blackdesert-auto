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

	// setAccPW(accConfig.Account, accConfig.Password)

	// pwd, time := load2FA(accConfig.FAkey)

	// join()
}
