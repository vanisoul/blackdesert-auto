package main

import (
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	// closeblack()
	// robotgo.MoveMouse(1297, 411-55)
	for {
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

	os.Exit(0)
}

func taskFeatures() {

}

func checkMainScreen() (ok bool) {
	return false
}
