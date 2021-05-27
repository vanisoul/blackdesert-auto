package main

import (
	"fmt"
	"strconv"

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
	// heatingConfig, err := LoadConfigheating()
	// if err != nil {
	// 	log.Errorf("cannot load config:", err)
	// 	return
	// }
	// println(FormulaNameArrayToString(heatingConfig.Method[0].Formula, " + "))
	// infoConfig, err := LoadConfigInfo()
	// if err != nil {
	// 	log.Errorf("cannot load config:", err)
	// }
	// bitmap := robotgo.CaptureScreen(infoConfig.GameScreenX, infoConfig.GameScreenY, infoConfig.GameScreenW, infoConfig.GameScreenH)
	// errorPngName := fmt.Sprintf("log/AAA.png")
	// robotgo.SaveBitmap(bitmap, errorPngName)
	for {
		leftMosue(0, 0)
		succ := checkMainScreen(false)
		if succ {
			taskFeatures()
		} else {
			taskLogin()
		}
	}
}

func taskLogin() (succ bool) {
	// joindesktop()
	setLog("taskLogin", "執行登入", "")

	errorScreen("img/winError.png")

	closeblack()
	robotgo.MoveMouse(0, 0)
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
	robotgo.Sleep(3)
	robotgo.MoveMouse(0, 0)
	robotgo.Sleep(57)
	count := 0
	for count < 4 {
		setLog("LoginTask", "等待登入成功畫面", strconv.Itoa(count))
		succCh := checkMainScreen(false)
		if succCh {
			succ = true
			return
		}
		robotgo.Sleep(40)
		count = count + 1
	}
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
	endStatus()
}

func checkMainScreen(gotomain ...bool) (succ bool) {
	count := 3
	for count > 0 {
		setLog("checkMainScreen", "進行畫面檢查", strconv.Itoa(count))
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
	if len(gotomain) == 0 {
		setLog("checkMainScreen", "檢查錯誤關閉並重啟遊戲", "")
		main()
	}
	return
}

func chooseRole(role int) {
	setLog("chooseRole", "選擇角色", strconv.Itoa(role))
	rolesucc, x, y := whilescreen("img/role.png")
	if rolesucc {
		y = y - 865
		leftMosue(x, y+((role-1)*100))
		robotgo.Sleep(1)
		leftMosueforimg("img/join.png")
	}
}

func setSafe(safepwd string) {
	setLog("setSafe", "輸入安全密碼", "")
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
		setLog("startGame", "從客戶端開啟遊戲", "")
		leftMosue(x, y)
	} else {
		succupdata, _, _ := whilescreen("img/clientUpdate.png", 5)
		if succupdata {
			startGame()
		}
	}
}

func selectGraphics(grap int) {
	setLog("selectGraphics", "選擇顯示卡", strconv.Itoa(grap))
	succsel, x, y := whilescreenMany(100, "img/selectGraphics.png", "img/selectGraphics2.png")
	if succsel {
		leftMosue(x, y+(40*grap))
		return
	}
}
