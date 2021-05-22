package main

import "github.com/labstack/gommon/log"

func main() {
	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
	}

	openGameClient(infoConfig.GamePath)
	// accConfig, err := LoadConfigAccount()
	// if err != nil {
	// 	log.Errorf("cannot load config:", err)
	// }

	// setAccPW(accConfig.Account, accConfig.Password)

	// pwd, time := load2FA(accConfig.FAkey)

	// join()
}
