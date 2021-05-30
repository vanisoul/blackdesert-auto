package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func additionalMatters() {
	setLog("additionalMatters", "", "")
	//精靈骰子
	elfDice()

	//領取物品
	getMoveItems()

	//搬運物品
	moveItems()

	//裝箱工人確認
	//請取名木箱/礦箱工人
	tradeGoods()
}

func tradeGoods() {
	tradeGoodsConfig, err := LoadConfigtradeGoods()
	if err != nil {
		log.Errorf("cannot load tradeGoodsConfig:", err)
		return
	}
	if !tradeGoodsConfig.Status {
		return
	}

	checkMainScreen()
	robotgo.KeyTap("m")
	mSucc, _, _ := whilescreenEasy("img/findMap.png")
	if !mSucc {
		return
	}

	ctrlA := func() {
		robotgo.KeyTap("a", "control")
	}
	enter := func() {
		robotgo.KeyTap("enter")
	}
	textLocationBeforeAfter(34, 113, tradeGoodsConfig.City, ctrlA, enter)
	robotgo.Sleep(1)
	interval := 28
	leftMosue(30, 257+interval*tradeGoodsConfig.Seq)
	robotgo.Sleep(3)
	robotgo.MoveMouse(960, 500)
	scrollup(5)
	leftMosue(960, 500)
	trSucc := leftMosueforimg("img/transport.png")
	if !trSucc {
		robotgo.KeyTap("m")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		return
	}
	robotgo.KeyTap("esc")
	boxSucc := true
	for boxSucc {
		for {
			if checkCityImg() {
				break
			}
			robotgo.KeyTap("esc")
		}
		bSucc, bx, by := whilescreen("img/boxing.png", 5)
		boxSucc = bSucc
		leftMosue(bx, by)
		if !boxSucc {
			return
		}

		makeFunc := func(typeImg string, img ...string) bool {
			mSucc := leftMosueforimg("img/make.png")
			if !mSucc {
				return false
			}
			robotgo.MoveMouse(649, 644)
			scrolldown(15)
			su, ix, iy := whilescreenManyEasy(3, insertStrToFilenameTailArr(img, "Task")...)
			if su {
				typeImg1 := fmt.Sprintf("img/%s.png", typeImg)
				typeImg2 := fmt.Sprintf("img/%s2.png", typeImg)
				_, tix, tiy := whilescreenMany(2, typeImg1, typeImg2)
				leftMosue(tix, tiy)
				robotgo.MoveMouse(0, 0)
				robotgo.Sleep(1)
				leftMosue(ix, iy)
				leftMosueforimg("img/selectFre.png")
				maxItem()
				ok := leftMosueforimg("img/goTask.png")
				if ok {
					return true
				}
			}
			return false
		}

		//礦
		if screenYesOrNoEasy("img/mineral.png") {
			if makeFunc("mineral-worker", tradeGoodsConfig.Mineral...) {
				setLog("tradeGoods", "礦物裝箱派遣成功", "")
			}
			continue
		}

		//木頭
		if screenYesOrNoEasy("img/wood.png") {
			if makeFunc("wood-worker", tradeGoodsConfig.Wood...) {
				setLog("tradeGoods", "木頭裝箱派遣成功", "")
			}
			continue
		}

		//不如預期
		setLog("tradeGoods", "加工場城鎮有不如預期房屋", "")
	}
}

func elfDice() {
	setLog("additionalMatters", "elfDice", "")
	checkMainScreen()
	robotgo.KeyTap("esc")
	eSucc, ex, ey := whilescreenMany(4, "img/elf.png", "img/elf2.png")
	if eSucc {
		setLog("elfDice", "found elfDice", "")
		leftMosue(ex, ey)
		robotgo.Sleep(1)
		leftMosueforimg("img/getElf.png")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
		robotgo.Sleep(1)
		robotgo.KeyTap("esc")
	}
	robotgo.KeyTap("esc")
}
