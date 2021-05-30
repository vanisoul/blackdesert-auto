package main

import (
	"github.com/go-vgo/robotgo"
)

func additionalMatters() {
	setLog("additionalMatters", "", "")
	//精靈骰子
	elfDice()

	//領取物品
	getMoveItems()

	//搬運物品
	moveItems()
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
