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
