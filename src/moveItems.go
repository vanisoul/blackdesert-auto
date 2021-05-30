package main

import (
	"fmt"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func getMoveItems() {
	moveItemsConfig, err := LoadConfigMoveItems()
	if err != nil {
		log.Errorf("cannot load moveItemsConfig:", err)
		return
	}
	if !moveItemsConfig.Status {
		return
	}

	for _, recy := range moveItemsConfig.Recycle {
		checkMainScreen()
		robotgo.KeyTap("m")
		mSucc, _, _ := whilescreenEasy("img/findMap.png")
		if !mSucc {
			continue
		}

		ctrlA := func() {
			robotgo.KeyTap("a", "control")
		}
		enter := func() {
			robotgo.KeyTap("enter")
		}
		textLocationBeforeAfter(34, 113, recy.City, ctrlA, enter)
		robotgo.Sleep(1)
		interval := 28
		leftMosue(30, 257+interval*recy.Seq)
		robotgo.Sleep(3)
		robotgo.MoveMouse(960, 500)
		scrollup(5)
		leftMosue(960, 500)
		trSucc := leftMosueforimg("img/transport.png")
		if !trSucc {
			robotgo.KeyTap("m")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			continue
		}
		trLogoSucc, _, _ := whilescreen("img/transportLogo.png")
		if !trLogoSucc {
			robotgo.KeyTap("m")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			continue
		}
		_, gx, gy := whilescreenMany(20, "img/getItems.png", "img/getItems2.png")
		leftMosue(gx, gy)

		robotgo.KeyTap("m")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
	}

}

func moveItems() {
	moveItemsConfig, err := LoadConfigMoveItems()
	if err != nil {
		log.Errorf("cannot load moveItemsConfig:", err)
		return
	}
	if !moveItemsConfig.Status {
		return
	}

	for _, move := range moveItemsConfig.Moves {
		setLog("moveItem", "開始移動", move.Source)
		checkMainScreen()
		robotgo.KeyTap("m")
		mSucc, _, _ := whilescreenEasy("img/findMap.png")
		if !mSucc {
			continue
		}

		ctrlA := func() {
			robotgo.KeyTap("a", "control")
		}
		enter := func() {
			robotgo.KeyTap("enter")
		}
		textLocationBeforeAfter(34, 113, move.Source, ctrlA, enter)
		robotgo.Sleep(1)
		interval := 28
		leftMosue(30, 257+interval*move.SourceSeq)
		robotgo.Sleep(3)
		robotgo.MoveMouse(960, 490)
		scrollup(5)
		leftMosue(960, 490)
		trSucc := leftMosueforimg("img/transport.png")
		if !trSucc {
			robotgo.KeyTap("m")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			continue
		}
		trLogoSucc, tlx, tly := whilescreen("img/transportLogo.png")
		if !trLogoSucc {
			robotgo.KeyTap("m")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			continue
		}
		leftMosue(tlx+340, tly+55)

		sdSucc := selectDes(move.Destination)
		if !sdSucc {
			robotgo.KeyTap("m")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			continue
		}
		stSucc := selectType(move.HaulType)
		if !stSucc {
			robotgo.KeyTap("m")
			robotgo.Sleep(1)
			robotgo.KeyTap("esc")
			continue
		}
		fullSucc := movePutAll(move.HaulType, move.Items...)
		sendItem(fullSucc, move.HaulType)
		robotgo.KeyTap("m")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
	}
}

func sendItem(fullSucc bool, HaulType string) {
	task := func() {
		r := saveIMG(true)
		setLog("sendItem", "發送物品", strconv.Itoa(r))
		leftMosueforimg("img/sendItem.png")
	}
	if HaulType == "trading" || fullSucc {
		task()
	}

}

func selectType(haulType string) (succ bool) {
	succ = false
	_, sx, sy := whilescreenEasy("img/selectCar.png")
	leftMosue(sx, sy)
	robotgo.Sleep(1)
	if haulType == "general" {
		succ = leftMosueforimg("img/generalCar.png")
	} else if haulType == "trading" {
		robotgo.MoveMouse(sx-20, sy+30)
		scrolldown(3)
		succ = leftMosueforimg("img/tradingCar.png")
	}
	return
}

func selectDes(des string) (succ bool) {
	succ = false
	desSucc, dx, dy := whilescreenEasy("img/destination.png")

	if !desSucc {
		return
	}
	leftMosue(dx, dy)
	robotgo.Sleep(1)
	count := 5
	for count > 0 {

		desFullName := fmt.Sprintf("img/%s.png", des)
		succ = leftMosueforimgCount(desFullName, 5)
		if succ {
			return
		}
		robotgo.MoveMouse(dx-20, dy+30)
		scrolldown(4)
		count = count + 1
	}
	return
}

func checkFullWeight(img string) (succ bool) {
	succ = false
	robotgo.MoveMouse(0, 0)
	robotgo.Sleep(1)
	bit := robotgo.CaptureScreen(1400, 250, 500, 600)
	defer robotgo.FreeBitmap(bit)
	if img == "煤炭.png" || img == "鋅礦石.png" {
		succ = !screenYesOrNoDotSelfimg(insertStrToFilenameTail(img, "full"), bit, 0.03, 2)
	} else {
		succ = screenYesOrNoDotSelfimg(insertStrToFilenameTail(img, "bank"), bit, 0.03, 2)
	}
	return
}

func movePutAll(haulType string, imgs ...string) (succ bool) {
	takeArticleSum := 0
	count := 2
	for count > 0 {
		for _, img := range imgs {
			succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "bank"), 6)
			if succright {
				robotgo.Sleep(1)
				maxItem()
				setLog("movePutAll", "加入貨運", img)
				takeArticleSum = takeArticleSum + 1
				if haulType == "general" {
					fuWSucc := checkFullWeight(img)
					if fuWSucc {
						setLog("movePutAll", "貨運已滿", "")
						succ = fuWSucc
						return
					}
				}

			}
			if takeArticleSum == len(imgs) {
				log.Info("takeArticleSum :", takeArticleSum)
				log.Info("len(imgs) :", len(imgs))
				return
			}
		}
		scrollmoveDown(6)
		count = count - 1
	}
	return
}

func scrollmoveDown(fre int) {
	_, x, y := whilescreen("img/ProcessingOK.png")
	log.Infof("scrollmoveDown x:%d, y:%d", x, y)
	robotgo.MoveMouse(x+41, y+103)
	scrolldown(fre)
}
