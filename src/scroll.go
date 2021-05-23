package main

import "github.com/go-vgo/robotgo"

func scrollBagDown(fre int) {
	_, x, y := whilescreen("img/bank_ok.png")
	robotgo.MoveMouse(x-3, y+151)
	scrolldown(fre)
}

func scrollBankDown(fre int) {
	_, x, y := whilescreen("img/bank_ok.png")
	robotgo.MoveMouse(x-148, y+131)
	scrolldown(fre)
}

func scrolldown(count int) {
	robotgo.Sleep(1)
	for count > 0 {
		robotgo.ScrollMouse(1, "down")
		robotgo.Sleep(1)
		count = count - 1
	}
}

func scrollup(count int) {
	robotgo.Sleep(1)
	for count > 0 {
		robotgo.ScrollMouse(1, "up")
		robotgo.Sleep(1)
		count = count - 1
	}
}
