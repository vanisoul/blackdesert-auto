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

		tksucc := takeRepoOne(beerConfig.Count, beerConfig.ArticlesTake...)
		if tksucc {
			setLog("beerTask", "從倉庫中取得其一恢復道具", strings.Join(beerConfig.ArticlesTake, ", "))
		}
		checkMainScreen()
		hot := strings.ToLower(beerConfig.HotKey)
		robotgo.KeyTap(hot)
		beerResucc := leftMosueforimg("img/beerRe.png")
		if beerResucc {
			setLog("beerTask", "開啟恢復勞工視窗", "")
		}
		beerEnsucc := leftMosueforimg("img/beerEnter.png")
		if beerEnsucc {
			setLog("beerTask", "使用恢復勞工道具", "")
		}
		robotgo.Sleep(3)
		restrsucc, x, y := whilescreenMany(20, "img/beerReStart.png", "img/beerReStart2.png")
		if restrsucc {
			leftMosue(x, y)
			setLog("beerTask", "勞工重新開始工作", "")
		}
		robotgo.Sleep(1)
		searchRepo()
		beerArtsucc := saveRepoOne(beerConfig.ArticlesSave...)
		if beerArtsucc {
			setLog("beerTask", "存放回倉庫其一恢復道具", strings.Join(beerConfig.ArticlesSave, ", "))
		}
	}
}
