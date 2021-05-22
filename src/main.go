package main

import (
	"github.com/labstack/gommon/log"
)

func main() {
	closeblack()
	// robotgo.MoveMouse(1080, 383)
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
	closeblack()

	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	opensucc := openGameClient(infoConfig.GamePath)
	if !opensucc {
		return
	}
	accConfig, err := LoadConfigAccount()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}

	setaccpwsuu := setAccPW(accConfig.Account, accConfig.Password)
	if !setaccpwsuu {
		return
	}
	// a := "a"
	// if accConfig.FAkey != "" {
	// 	pwd, time := load2FA(accConfig.FAkey)
	// }

}

func taskFeatures() {

}

func checkMainScreen() (ok bool) {
	return false
}
