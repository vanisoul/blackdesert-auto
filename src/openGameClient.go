package main

import (
	"os/exec"
	"path/filepath"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func openGameClient(path string) (succ bool) {
	cmd := exec.Command(filepath.Join(path, "BlackDesertLauncher.exe"))
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error:", err)
		return
	}

	succok, _, _ := whilescreen("img/gameclientok.png")
	if succok {
		succ = true
		return
	}

	saveIMG()
	rsucc, _, _ := whilescreen("img/repair.png")
	saveIMG()
	if rsucc {
		saveIMG()
		robotgo.Sleep(7200)
		saveIMG()
		main()
	} else {
		saveIMG()
		main()
	}
	return
}
