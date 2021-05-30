package main

import "github.com/go-vgo/robotgo"

func textLocationBefore(x int, y int, text string, args ...func()) (succ bool) {

	leftMosue(x, y)

	if len(args) == 1 {
		fc := args[0]
		fc()
	}
	robotgo.Sleep(1)
	robotgo.TypeStr(text, 0.1)

	succ = true
	return

}

func textLocationAfter(x int, y int, text string, args ...func()) (succ bool) {

	leftMosue(x, y)
	robotgo.TypeStr(text, 0.1)
	robotgo.Sleep(1)
	if len(args) == 1 {
		fc := args[0]
		fc()
	}

	succ = true
	return

}

func textLocationBeforeAfter(x int, y int, text string, args ...func()) (succ bool) {

	leftMosue(x, y)

	if len(args) >= 1 && args[0] != nil {
		fc := args[0]
		fc()
	}

	robotgo.Sleep(1)
	robotgo.TypeStr(text, 0.1)
	robotgo.Sleep(1)

	if len(args) >= 2 && args[1] != nil {
		fc := args[1]
		fc()
	}

	succ = true
	return

}
