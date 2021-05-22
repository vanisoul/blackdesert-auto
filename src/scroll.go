package main

import "github.com/go-vgo/robotgo"

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
