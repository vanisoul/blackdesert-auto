package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func main() {
	// robotgo.DragMouse()

	// fxx, fyy := robotgo.GetMousePos()
	// fmt.Println("FindBitmap------", fxx, fyy)
	// _, fx, fy := whilescreen("img/ProcessingOK.png")
	// fmt.Println("FindBitmap------", fx, fy)
	// proing, _, _ := whilescreenEasy("img/Processeding_1.png")
	// fmt.Println(proing)

	// heatingConfig, err := LoadConfigheating()
	// if err != nil {
	// 	log.Errorf("cannot load config:", err)
	// 	return
	// }
	// runTask("Heating", heatingConfig.Method)

	for {
		succ := checkMainScreen(false)
		if succ {
			taskFeatures()
		} else {
			loginsucc := taskLogin()
			if loginsucc {
				searchRepo()
				infoConfig, err := LoadConfigInfo()
				if err != nil {
					log.Errorf("cannot load config:", err)
					return
				}
				saveRepoAll(infoConfig.ClearBag...)
			}
		}
	}
}

func taskLogin() (succ bool) {
	// joindesktop()
	closeblack()

	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
		succ = false
		return
	}
	opensucc := openGameClient(infoConfig.GamePath)
	if !opensucc {
		log.Errorf("openGameClient")
		succ = false
		return
	}
	accConfig, err := LoadConfigAccount()
	if err != nil {
		log.Errorf("cannot load config:", err)
		succ = false
		return
	}

	setaccpwsuu := setAccPW(accConfig.Account, accConfig.Password)
	if !setaccpwsuu {
		log.Errorf("setaccpw")
		succ = false
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

	gsscc, gx, gy := whilescreen("img/gameStart.png", 200)
	if gsscc {
		leftMosue(gx, gy)
	} else {
		succ = false
		return
	}
	leftMosueforimg("img/agree.png")

	setSafe(accConfig.Safe)
	leftMosueforimg("img/randomLogin.png")

	chooseRole(infoConfig.Role)
	robotgo.Sleep(10)
	robotgo.MoveMouse(0, 0)
	robotgo.Sleep(20)
	succ = true
	return
}

func taskFeatures() {
	//勞工恢復體力
	beerTask()
	checkMainScreen()
	// 加熱
	heatingTask()
	// checkMainScreen()
	//勞工恢復體力
	// beerTask()
	// checkMainScreen()
	// 砍材
	// chopWoodTask()
	// checkMainScreen()
	// 料理
	// checkMainScreen()
	// 煉金
}

func checkMainScreen(gotomain ...bool) (succ bool) {
	count := 3
	for count > 0 {
		robotgo.KeyTap("esc")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		succ, _, _ = whilescreen("img/esc.png", 2)
		if succ {
			robotgo.MoveMouse(0, 0)
			robotgo.KeyTap("esc")
			robotgo.Sleep(1)
			return
		}
		count = count - 1
	}
	closeblack()
	if gotomain[0] {
		main()
	}
	return
}

func chooseRole(role int) {
	rolesucc, x, y := whilescreen("img/role.png")
	if rolesucc {
		y = y - 865
		leftMosue(x, y+((role-1)*100))
		robotgo.Sleep(1)
		leftMosueforimg("img/join.png")
	}
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
	succsel, x, y := whilescreenMany(100, "img/selectGraphics.png", "img/selectGraphics2.png")
	if succsel {
		leftMosue(x, y+(40*grap))
		return
	}
}
