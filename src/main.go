package main

import (
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func main() {
	// robotgo.DragMouse()

	// fx, fy := robotgo.GetMousePos()
	// fmt.Println("FindBitmap------", fx, fy)

	// robotgo.DragMouse(fx, fy)
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

	leftMosueforimg("img/gameStart.png")
	leftMosueforimg("img/agree.png")

	setSafe(accConfig.Safe)

	os.Exit(0)
}

func taskFeatures() {

}

func checkMainScreen() (ok bool) {
	return false
}

func setSafe(safepwd string) {
	for _, pwd := range safepwd {
		robotgo.Sleep(2)
		img := fmt.Sprintf("img/%s.png", string(pwd))
		leftMosueforimg(img)
	}
	leftMosueforimg("img/setSafePW.png")

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
	count := 100
	for count > 0 {
		succsel, x, y := whilescreen("img/selectGraphics.png", 1)
		if succsel {
			leftMosue(x, y+(40*grap))
			return
		}
		succsel2, x, y := whilescreen("img/selectGraphics2.png", 1)
		if succsel2 {
			leftMosue(x, y+(40*grap))
			return
		}
		robotgo.Sleep(1)
		count = count - 1
	}
}
