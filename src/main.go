package main

import (
	"github.com/labstack/gommon/log"
)

func main() {

	for false {
		succ := checkMainScreen()
		if succ {
			taskFeatures()
		} else {
			taskLogin()
		}
	}
}

func taskLogin() {
	// joindesktop()
	// closeblack()

	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
	}
	openGameClient(infoConfig.GamePath)
	accConfig, err := LoadConfigAccount()
	if err != nil {
		log.Errorf("cannot load config:", err)
	}

	setAccPW(accConfig.Account, accConfig.Password)

	// if accConfig.FAkey != "" {
	// 	pwd, time := load2FA(accConfig.FAkey)
	// }

}

func taskFeatures() {

}

func checkMainScreen() (ok bool) {
	return false
}
