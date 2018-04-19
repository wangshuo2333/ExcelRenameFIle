package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	var filePre string
	if len(os.Args) > 1 {
		filePre = os.Args[1]
	} else {
		filePre, _ = os.Getwd()
	}
	filePre = filePre + "/"
	xlFile, err := xlsx.OpenFile(filePre + "file.xlsx")
	CheckErr(err)
	oldPathName := ""
	newPathName := ""
	for _, sheet := range xlFile.Sheets {
		for rowNum, row := range sheet.Rows {
			for columnNum, cell := range row.Cells {
				if rowNum == 0 {
					continue
				}
				if (columnNum == 0) && (cell.Type() == 0) {
					oldPathName = cell.String()
				}

				if (columnNum == 1) && (cell.Type() == 0) {
					newPathName = cell.String()
				}
				continue
			}
			if oldPathName != "" && newPathName != "" {
				err = os.Rename(filePre+oldPathName, filePre+newPathName)
				fmt.Println(err)
			}
			oldPathName = ""
			newPathName = ""
		}
	}

}
