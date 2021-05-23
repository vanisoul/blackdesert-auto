package main

import (
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

func beerTask() {
	beerConfig, err := LoadConfigbeer()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	if beerConfig.Status {
		searchRepo()

		takeRepoOne(beerConfig.Count, beerConfig.ArticlesTake...)
		checkMainScreen()
		hot := strings.ToLower(beerConfig.HotKey)
		robotgo.KeyTap(hot)
		leftMosueforimg("img/beerRe.png")
		leftMosueforimg("img/beerEnter.png")
		robotgo.Sleep(3)
		restrsucc, x, y := whilescreenMany(20, "img/beerReStart.png", "img/beerReStart2.png")
		if restrsucc {
			leftMosue(x, y)
		}
		robotgo.Sleep(1)
		searchRepo()
		saveRepoOne(beerConfig.ArticlesSave...)
	}
}
