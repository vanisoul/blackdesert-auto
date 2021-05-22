package main

import "github.com/go-vgo/robotgo"

func setAccPW(acc string, pw string) {
	ctrlA := func() {
		robotgo.KeyTap("a", "control")
	}
	clickLocation("login.img", 0, 20, acc, ctrlA)
	clickLocation("login.img", 0, 10, pw)
}
