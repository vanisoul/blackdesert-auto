package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func errorScreen(file string) {
	errSucc, errx, erry := whilescreen(file, 3)
	if errSucc {

		infoConfig, err := LoadConfigInfo()
		if err != nil {
			log.Errorf("cannot load config:", err)
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
		bitmap := robotgo.CaptureScreen(infoConfig.GameScreenX, infoConfig.GameScreenY, infoConfig.GameScreenW, infoConfig.GameScreenH)
		errorPngName := fmt.Sprintf("log/%d.png", r)
		robotgo.SaveBitmap(bitmap, errorPngName)
		setLog("taskLogin", "windows 有錯誤視窗", strconv.Itoa(r))
		leftMosue(errx, erry)
	}
}

//傳入true 就不會設定預設log
func saveIMG(args ...bool) (r int) {
	infoConfig, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load config:", err)
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	bitmap := robotgo.CaptureScreen(infoConfig.GameScreenX, infoConfig.GameScreenY, infoConfig.GameScreenW, infoConfig.GameScreenH)
	errorPngName := fmt.Sprintf("log/%d.png", r)
	robotgo.SaveBitmap(bitmap, errorPngName)
	if len(args) == 1 && args[0] {
		return
	}
	setLog("taskLogin", "被迫關閉遊戲", strconv.Itoa(r))
	return
}
