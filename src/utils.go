package main

import (
	"fmt"
	"path"
	"strings"
)

func insertStrToFilenameTailArr(strs []string, taillStr string) (res []string) {
	newstrs := []string{}
	for _, str := range strs {

		ext := path.Ext(str)
		name := strings.ReplaceAll(str, ext, "")
		strAnswer := fmt.Sprintf("%s_%s%s", name, taillStr, ext)
		newstrs = append(newstrs, strAnswer)
	}
	res = newstrs
	return
}

func insertStrToFilenameTail(str string, taillStr string) (res string) {

	ext := path.Ext(str)
	name := strings.ReplaceAll(str, ext, "")
	res = fmt.Sprintf("%s_%s%s", name, taillStr, ext)

	return
}
