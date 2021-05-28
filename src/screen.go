package main

import (
	"fmt"
	"image"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

// func screen() {
// 	configInfo, err := LoadConfigInfo()
// 	if err != nil {
// 		log.Errorf("cannot load configInfo:", err)
// 	}
// 	x, err := strconv.Atoi(configInfo.GameScreenX)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	y, err := strconv.Atoi(configInfo.GameScreenY)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	w, err := strconv.Atoi(configInfo.GameScreenW)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	h, err := strconv.Atoi(configInfo.GameScreenH)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	screenXY(x, y, w, h)
// }

// func screenDebug() {
// 	sx := os.Args[2]
// 	sy := os.Args[3]
// 	sw := os.Args[4]
// 	sh := os.Args[5]
// 	x, err := strconv.Atoi(sx)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	y, err := strconv.Atoi(sy)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	w, err := strconv.Atoi(sw)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	h, err := strconv.Atoi(sh)
// 	if err != nil {
// 		log.Errorf("Error strconv:", err)
// 		return
// 	}
// 	screenXY(x, y, w, h)
// }

// func screenXY(x int, y int, w int, h int) {
// 	bitmap := robotgo.CaptureScreen(x, y, w, h)
// 	defer robotgo.FreeBitmap(bitmap)
// 	fmt.Println("...", bitmap)
// 	fx, fy := robotgo.FindBitmap(bitmap)
// 	fmt.Println("FindBitmap------ ", fx, fy)
// 	robotgo.SaveBitmap(bitmap, "tmp.png")
// }

func whilescreen(pngName string, jcount ...int) (succ bool, x int, y int) {

	file, _ := os.Open(pngName)
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Errorf("load pngName Error:%s, pngName: %s", err, pngName)
		file.Close()
		return
	}

	file.Close()

	imgh := c.Height
	imgw := c.Width
	CopyFile(pngName, "tmp.png")
	bit_map := robotgo.OpenBitmap("tmp.png")

	defer robotgo.FreeBitmap(bit_map)

	count := 20
	if len(jcount) == 1 {
		count = int(jcount[0])
	}
	log.Info("screen", pngName)
	log.Info("count", count)
	for {
		robotgo.Sleep(1)

		fx, fy := robotgo.FindBitmap(bit_map)

		fmt.Println("FindBitmap------", fx, fy)
		if fx != -1 && fy != -1 {
			succ = true
			x = fx + (imgw / 2)
			y = fy + (imgh / 2)
			return
		}
		count = count - 1
		if count == 0 {
			succ = false
			x = -1
			y = -1
			return
		}
	}

}

func screenYesOrNoEasy(pngName string, jcount ...int) (succ bool) {
	CopyFile(pngName, "tmp.png")
	bit_map := robotgo.OpenBitmap("tmp.png")
	defer robotgo.FreeBitmap(bit_map)
	count := 20
	if len(jcount) == 1 {
		count = int(jcount[0])
	}
	log.Info("screen", pngName)
	log.Info("count", count)
	for {
		robotgo.Sleep(1)

		fx, fy := robotgo.FindBitmap(bit_map, nil, 0.1)

		fmt.Println("FindBitmap------", fx, fy)
		if fx != -1 && fy != -1 {
			succ = true
			return
		}
		count = count - 1
		if count == 0 {
			succ = false
			return
		}
	}
}

func whilescreenEasy(pngName string, jcount ...int) (succ bool, x int, y int) {

	file, _ := os.Open(pngName)
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Errorf("load pngName Error:%s, pngName: %s", err, pngName)
		file.Close()
		return
	}
	file.Close()

	imgh := c.Height
	imgw := c.Width
	CopyFile(pngName, "tmp.png")
	bit_map := robotgo.OpenBitmap("tmp.png")

	defer robotgo.FreeBitmap(bit_map)

	count := 20
	if len(jcount) == 1 {
		count = int(jcount[0])
	}
	log.Info("screen", pngName)
	log.Info("count", count)
	for {
		robotgo.Sleep(1)

		fx, fy := robotgo.FindBitmap(bit_map, nil, 0.1)

		fmt.Println("FindBitmap------", fx, fy)
		if fx != -1 && fy != -1 {
			succ = true
			x = fx + (imgw / 2)
			y = fy + (imgh / 2)
			return
		}
		count = count - 1
		if count == 0 {
			succ = false
			x = -1
			y = -1
			return
		}
	}

}

func whilescreenMany(count int, imgNames ...string) (succ bool, resx int, resy int) {

	for count > 0 {
		for _, imgName := range imgNames {
			succScr, x, y := whilescreen(imgName, 1)
			if succScr {
				succ = true
				resx = x
				resy = y
				return
			}
		}
		robotgo.Sleep(1)
		count = count - 1
	}
	succ = false
	resx = -1
	resy = -1
	return
}

func whilescreenManyEasy(count int, imgNames ...string) (succ bool, resx int, resy int) {

	for count > 0 {
		for _, imgName := range imgNames {
			succScr, x, y := whilescreenEasy(imgName, 1)
			if succScr {
				succ = true
				resx = x
				resy = y
				return
			}
		}
		robotgo.Sleep(1)
		count = count - 1
	}
	succ = false
	resx = -1
	resy = -1
	return
}
