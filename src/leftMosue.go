package main

import "github.com/go-vgo/robotgo"

func leftMosue(x int, y int) {
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`down`, `left`)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `left`)
}

func leftMosueforimg(img string) (succ bool) {
	succscr, x, y := whilescreen(img)
	if succscr {
		robotgo.MoveMouse(x, y)
		robotgo.Sleep(1)
		robotgo.MouseToggle(`down`, `left`)
		robotgo.Sleep(1)
		robotgo.MouseToggle(`up`, `left`)
		succ = true
		return
	} else {
		succ = false
		return
	}
}
