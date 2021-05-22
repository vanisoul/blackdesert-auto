package main

import "github.com/go-vgo/robotgo"

func setAccPW(acc string, pw string) (succ bool) {
	ctrlA := func() {
		robotgo.KeyTap("a", "control")
	}
	accsucc := clickLocation("img/login.png", 20, 20, acc, ctrlA)
	if !accsucc {
		return
	}
	pwsucc := clickLocation("img/login.png", 20, 10, pw)
	if !pwsucc {
		return
	}
	succ = true
	return
}
