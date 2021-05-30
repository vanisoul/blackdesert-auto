package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func insertStrToFilenameTailArr(strs []string, taillStr string) (res []string) {
	newstrs := []string{}
	for _, str := range strs {

		ext := path.Ext(str)
		name := strings.ReplaceAll(str, ext, "")
		fullName := path.Join("img", "items", name)
		strAnswer := fmt.Sprintf("%s_%s%s", fullName, taillStr, ext)
		newstrs = append(newstrs, strAnswer)
	}
	res = newstrs
	return
}

func setArmsPathArr(strs []string) (res []string) {
	newstrs := []string{}
	for _, str := range strs {

		fullName := path.Join("img", "arms", str)
		newstrs = append(newstrs, fullName)
	}
	res = newstrs
	return
}

func setPearlArmsPathArr(strs []string) (res []string) {
	newstrs := []string{}
	for _, str := range strs {

		fullName := path.Join("img", "pearlArms", str)
		newstrs = append(newstrs, fullName)
	}
	res = newstrs
	return
}

//將檔案名稱尾插入一段字
func insertStrToFilenameTail(str string, taillStr string) (res string) {

	ext := path.Ext(str)
	name := strings.ReplaceAll(str, ext, "")
	fullName := path.Join("img", "items", name)
	res = fmt.Sprintf("%s_%s%s", fullName, taillStr, ext)
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

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func FormulaNameArrayToString(fml []formula, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(fml), " ", delim, -1), "[]")
}

func maxItem() {
	robotgo.Sleep(1)
	robotgo.KeyTap("f")
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
	robotgo.Sleep(1)
}
