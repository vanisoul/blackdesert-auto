package main

import "github.com/go-vgo/robotgo"

func leftMosue(x int, y int) {
	robotgo.MoveMouse(x, y)
	robotgo.Sleep(1)
	robotgo.MouseClick("left", true)
}
