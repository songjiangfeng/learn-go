package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

//获取指定目录下的所有文件和目录
func main() {

	excelFileName := "/Users/songjiangfeng/Desktop/products.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
				break
			}
		}
	}
}
