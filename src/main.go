package main

import (
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	// closeblack()
	// robotgo.MoveMouse(954, 416+40+40)
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
		log.Errorf("openGameClient")
		return
	}
	accConfig, err := LoadConfigAccount()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}

	setaccpwsuu := setAccPW(accConfig.Account, accConfig.Password)
	if !setaccpwsuu {
		log.Errorf("setaccpw")
		return
	}

	if accConfig.FAkey != "" {
		pwd, _ := load2FA(accConfig.FAkey)
		setFA(pwd)
	}
	startGame()
	if infoConfig.SelectGraphics != 0 {
		selectGraphics(infoConfig.SelectGraphics)
	}
	os.Exit(0)
}

func taskFeatures() {

}

func checkMainScreen() (ok bool) {
	return false
}

func startGame() {
	succstart, x, y := whilescreen("img/clientStart.png")
	if succstart {
		leftMosue(x, y)
	} else {
		succupdata, _, _ := whilescreen("img/clientUpdate.png", 5)
		if succupdata {
			startGame()
		}
	}
}

func selectGraphics(grap int) {
	succsel, x, y := whilescreen("img/selectGraphics.png")
	if succsel {
		leftMosue(x, y+(40*grap))
	}

}
