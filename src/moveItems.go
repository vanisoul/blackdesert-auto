package main

import (
	"fmt"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

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
		robotgo.MoveMouse(960, 500)
		scrollup(5)
		leftMosue(960, 500)
		trSucc := leftMosueforimg("img/transport.png")
		if !trSucc {
			continue
		}
		trLogoSucc, tlx, tly := whilescreen("img/transportLogo.png")
		if !trLogoSucc {
			continue
		}
		leftMosue(tlx+340, tly+55)

		sdSucc := selectDes(move.Destination)
		if !sdSucc {
			continue
		}
		stSucc := selectType(move.HaulType)
		if !stSucc {
			return
		}
		fullSucc := movePutAll(move.Items...)
		sendItem(fullSucc, move.HaulType)

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
	succ, _, _ = whilescreenEasy(insertStrToFilenameTail(img, "bank"), 2)
	return
}

func movePutAll(imgs ...string) (succ bool) {
	takeArticleSum := 0
	count := 2
	for count > 0 {
		for _, img := range imgs {
			succright := rightMosueforimgEasy(insertStrToFilenameTail(img, "bank"), 6)
			if succright {
				robotgo.Sleep(1)
				maxItem()
				takeArticleSum = takeArticleSum + 1
				fuWSucc := checkFullWeight(img)
				setLog("movePutAll", "加入貨運", img)
				if fuWSucc {
					setLog("movePutAll", "貨運已滿", "")
					succ = fuWSucc
					return
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
