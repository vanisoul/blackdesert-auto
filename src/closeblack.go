package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func closeblack() {
	fpid, err := robotgo.FindIds("BlackDesert")
	if err == nil {
		fmt.Println("pids... ", fpid)
		if len(fpid) > 0 {
			for _, pid := range fpid {
				robotgo.ActivePID(pid)
				robotgo.Kill(pid)
			}

		}
	}

}
