package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func screen() {
	configInfo, err := LoadConfigInfo()
	if err != nil {
		log.Errorf("cannot load configInfo:", err)
	}
	x, err := strconv.Atoi(configInfo.GameScreenX)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	y, err := strconv.Atoi(configInfo.GameScreenY)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	w, err := strconv.Atoi(configInfo.GameScreenW)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	h, err := strconv.Atoi(configInfo.GameScreenH)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	screenXY(x, y, w, h)
}

func screenDebug() {
	sx := os.Args[2]
	sy := os.Args[3]
	sw := os.Args[4]
	sh := os.Args[5]
	x, err := strconv.Atoi(sx)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	y, err := strconv.Atoi(sy)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	w, err := strconv.Atoi(sw)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	h, err := strconv.Atoi(sh)
	if err != nil {
		log.Errorf("Error strconv:", err)
		return
	}
	screenXY(x, y, w, h)
}

func screenXY(x int, y int, w int, h int) {
	bitmap := robotgo.CaptureScreen(x, y, w, h)
	defer robotgo.FreeBitmap(bitmap)

	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------ ", fx, fy)

	robotgo.SaveBitmap(bitmap, "tmp.png")
}

func whilescreen(pngName string) {

}
