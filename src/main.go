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

	// beerConfig, err := LoadConfigbeer()
	// if err != nil {
	// 	log.Errorf("cannot load config:", err)
	// 	return
	// }
	// saveRepoAll(beerConfig.ArticlesSave...)
	// bit := robotgo.OpenBitmap("img/3000_VoL.png")
	// fx, fy := robotgo.FindBitmap(bit, nil, 0.1)
	// fmt.Println("FindBitmap------", fx, fy)
	// robotgo.MoveMouse(fx, fy)
	// bit := robotgo.OpenBitmap("img/ProcessedClothing.png")
	// co := 20
	// for co > 0 {
	// 	fx, fy := robotgo.FindBitmap(bit, nil, 0.1)
	// 	fmt.Println("FindBitmap------", fx, fy)
	// 	co = co - 1
	// }
	checkMainScreen()
	heatingTask()
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

	gsscc, gx, gy := whilescreen("img/gameStart.png", 200)
	if gsscc {
		leftMosue(gx, gy)
	} else {
		return
	}
	leftMosueforimg("img/agree.png")

	setSafe(accConfig.Safe)
	leftMosueforimg("img/randomLogin.png")

	chooseRole(infoConfig.Role)
	robotgo.Sleep(20)
}

func taskFeatures() {
	//勞工恢復體力
	beerTask()
	checkMainScreen()
	// 加熱
	heatingTask()
	// checkMainScreen()
	// 砍材
	// checkMainScreen()
	// 料理
	// checkMainScreen()
	// 煉金
}

func checkMainScreen() (succ bool) {
	count := 3
	for count > 0 {
		robotgo.KeyTap("esc")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		succ, _, _ = whilescreen("img/esc.png", 2)
		if succ {
			robotgo.KeyTap("esc")
			robotgo.Sleep(2)
			return
		}
		count = count - 1
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
