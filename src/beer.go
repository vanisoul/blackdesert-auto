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
		leftMosueforimg("img/beerReStart.png")
		robotgo.Sleep(1)
		searchRepo()
		saveRepoOne(beerConfig.ArticlesSave...)
	}
}
