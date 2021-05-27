package main

import "github.com/go-vgo/robotgo"

func setAccPW(acc string, pw string) (succ bool) {
	setLog("setAccPW", "輸入帳號密碼", "")
	ctrlA := func() {
		robotgo.KeyTap("a", "control")
	}
	succsc, imgx, imgy := whilescreen("img/login.png")
	if succsc {
		accsucc := textLocation(imgx, imgy-105, acc, ctrlA)
		if !accsucc {
			return
		}
		pwsucc := textLocation(imgx, imgy-55, pw)
		if !pwsucc {
			return
		}
		leftMosue(imgx, imgy)
		succ = true
	} else {
		succ = false
	}
	return
}
