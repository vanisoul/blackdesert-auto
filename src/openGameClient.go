package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func openGameClient(path string) {
	cmd := exec.Command(filepath.Join(path, "BlackDesertLauncher.exe"))
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
