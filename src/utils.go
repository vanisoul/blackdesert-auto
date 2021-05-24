package main

import (
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"
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

//將檔案名稱尾插入一段字
func insertStrToFilenameTail(str string, taillStr string) (res string) {

	ext := path.Ext(str)
	name := strings.ReplaceAll(str, ext, "")
	res = fmt.Sprintf("%s_%s%s", name, taillStr, ext)

	return
}

//生成count個[start,end)結束的不重複的隨機數
func generateRandomNumber(start int, end int, count int) []int {
	//範圍檢查
	if end < start || (end-start) < count {
		return nil
	}
	//存放結果的slice
	nums := make([]int, 0)
	//隨機數生成器,加入時間戳保證每次生成的隨機數不一樣
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成隨機數
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
