package main

import "github.com/go-vgo/robotgo"

func textLocation(x int, y int, text string, args ...func()) (succ bool) {

	leftMosue(x, y)

	if len(args) == 1 {
		fc := args[0]
		fc()
	}
	robotgo.TypeStr(text, 0.1)

	succ = true
	return

}
