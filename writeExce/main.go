package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
)

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

//获取绝对路径
func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

//遍历目标地址所有文件
func PrintFilesName(path string) {
	fullPath := GetFullPath(path)
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	CheckErr(err)
	row := sheet.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "原始文件名"
	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		name := fi.Name()
		row := sheet.AddRow()
		row.SetHeightCM(1)
		cell := row.AddCell()
		cell.Value = name
		err1 := file.Save("file.xlsx")
		CheckErr(err1)
		return nil
	})
}

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path, _ = os.Getwd()
	}
	PrintFilesName(path)

	fmt.Println("done!")
}
