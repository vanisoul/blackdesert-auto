package main

import "github.com/go-vgo/robotgo"

func additionalMatters() {

	//精靈骰子
	elfDice()

	//搬運物品
	moveItems()
}

func elfDice() {
	checkMainScreen()
	robotgo.KeyTap("esc")
	eSucc, ex, ey := whilescreen("img/elf.png")
	if eSucc {
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

func moveItems() {

}
