package main

import "github.com/go-vgo/robotgo"

func leftMosue(x int, y int) {
	robotgo.Sleep(1)
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`down`, `left`)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `left`)
}

func leftMosueforimg(img string) (succ bool) {
	succscr, x, y := whilescreen(img)
	if succscr {
		robotgo.Sleep(1)
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

func rightMosue(x int, y int) {
	robotgo.Sleep(1)
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`down`, `right`)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `right`)
}

func rightMosueforimg(img string, count int) (succ bool) {
	succscr, x, y := whilescreen(img, count)
	if succscr {
		rightMosueforimgFunc(succscr, x, y)
		succ = true
	} else {
		succ = false
	}
	return
}

func rightMosueforimgEasyAll(imgs []string, count int) (succ bool) {
	for _, img := range imgs {
		succscr, x, y := whilescreenEasy(img, count)
		if succscr {
			rightMosueforimgFunc(succscr, x, y)
		} else {
			succ = false
		}
	}
	succ = true
	return
}

func rightMosueforimgEasy(img string, count int) (succ bool) {

	succscr, x, y := whilescreenEasy(img, count)
	succ = rightMosueforimgFunc(succscr, x, y)
	return
}

func rightMosueforimgFunc(succsrc bool, x int, y int) (succ bool) {
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`down`, `right`)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `right`)
	succ = true
	return
}
