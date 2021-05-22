package main

import "github.com/go-vgo/robotgo"

func clickLocation(imgName string, x int, y int, text string, args ...func()) (succ bool) {

	succ, imgx, imgy := whilescreen(imgName)

	if succ {
		leftMosue(imgx+x, imgy+y)
	} else {
		succ = false
	}

	if len(args) == 1 {
		fc := args[0]
		fc()
	}
	robotgo.TypeStr(text, 0.1)

	succ = true
	return

}
