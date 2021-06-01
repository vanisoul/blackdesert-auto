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

	rsucc, _, _ := whilescreenMany(5, "img/repair.png", "img/repair2.png")
	saveIMG()
	if rsucc {
		robotgo.Sleep(7200)
		main()
	}

	succok, _, _ := whilescreen("img/gameclientok.png")
	if succok {
		succ = true
		return
	}

	return
}
