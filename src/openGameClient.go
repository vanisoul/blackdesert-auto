package main

import (
	"os/exec"
	"path/filepath"

	"github.com/labstack/gommon/log"
)

func openGameClient(path string) {
	cmd := exec.Command(filepath.Join(path, "BlackDesertLauncher.exe"))
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error:", err)
		return
	}

	whilescreen("img/gameclientok.png")
}
