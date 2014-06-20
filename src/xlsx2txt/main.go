package main

import (
	"fmt"
	"github.com/jteeuwen/go-pkg-optarg"
	"github.com/tealeg/xlsx"
)

func main() {
	optarg.Add("c", "column-header", "This row is used as column name", 1)
	optarg.Add("f", "flag", "If a cell in this row is empty, these column is ignored", 1)
	optarg.Add("i", "id", "Cell value in this column is used as id of row", 1)
	optarg.Add("s", "sheet-filter", "Worksheet whose name starts with this text is ignored", "_")
	optarg.Add("h", "help", "Displays this help.", false)

	columnHeader := 1
	id := 1
	flag := 1
	sheetFilter := "_"

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "h":
			optarg.Usage()
			return
		case "c":
			columnHeader = opt.Int()
		case "f":
			flag = opt.Int()
		case "i":
			id = opt.Int()
		case "s":
			sheetFilter = opt.String()
		}
	}
	files := optarg.Remainder
	switch len(files) {
	case 0:
		fmt.Println("It needs file name as parameter to run, but no file specified\n")
		optarg.Usage()
	case 1:
		workbook, err := xlsx.OpenFile(files[0])
		if err != nil {
			fmt.Println(err)
			fmt.Printf("can't parse xlsx file: %s", files[0])
			return
		}
		WalkSheet(workbook, columnHeader, flag, id, sheetFilter)
	default:
		fmt.Printf("It needs one file name as parameter to run but %d files are passed\n\n", len(files))
		optarg.Usage()
	}
}
