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
		leftMosueforimgFunc(x, y)
		succ = true
		return
	} else {
		succ = false
		return
	}
}

func leftMosueforimgEasy(img string) (succ bool) {
	succscr, x, y := whilescreenEasy(img)
	if succscr {
		leftMosueforimgFunc(x, y)
		succ = true
		return
	} else {
		succ = false
		return
	}
}

func leftMosueforimgFunc(x int, y int) (succ bool) {
	robotgo.Sleep(1)
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`down`, `left`)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `left`)
	succ = true
	return
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
		rightMosueforimgFunc(x, y)
		succ = true
	} else {
		succ = false
	}
	return
}

func rightMosueforimgEasyAll(imgs []string, count int) (succ bool, succfre int) {
	tmpsucc := 0
	for _, img := range imgs {
		succscr, x, y := whilescreenEasy(img, count)
		if succscr {
			rightMosueforimgFunc(x, y)
			setLog("rightMosueforimgEasyAll", "成功點擊右鍵", img)
			tmpsucc = tmpsucc + 1
		} else {
			succ = false
		}
	}
	succfre = tmpsucc
	succ = true
	return
}

func rightMosueforimgEasy(img string, count int) (succ bool) {

	succscr, x, y := whilescreenEasy(img, count)
	if succscr {
		succ = rightMosueforimgFunc(x, y)
	} else {
		succ = false
	}

	return
}

func rightMosueforimgFunc(x int, y int) (succ bool) {
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`down`, `right`)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `right`)
	robotgo.Sleep(1)
	succ = true
	return
}
