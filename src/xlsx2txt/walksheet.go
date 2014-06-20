package main

import (
	"fmt"
	"strings"
	"github.com/tealeg/xlsx"
)

func WalkSheet(workbook *xlsx.File, columnHeader, flag, id int, sheetFilter string) {
	for _, worksheet := range workbook.Sheets {
		sheetname := worksheet.Name
		if strings.HasPrefix(sheetname, sheetFilter) {
			continue
		}
		var flags map[int]bool
		var headers map[int]string
		for i, row := range worksheet.Rows {
			processed := false
			if i == columnHeader - 1 {
				headers = ReadColumnHeaderRow(row)
				processed = true
			}
			if i == flag - 1 {
				flags = ReadFlagRow(row)
				processed = true
			}
			if (!processed && len(flags) > 0 && len(headers) > 0) {
				DumpRow(row, sheetname, flags, headers, id)
			}
		}
	}
}

func ReadColumnHeaderRow(row *xlsx.Row) map[int]string {
	var result = make(map[int]string)
	for i, cell := range row.Cells {
		if len(cell.String()) != 0 {
			result[i] = cell.String()
		}
	}
	return result
}

func ReadFlagRow(row *xlsx.Row) map[int]bool {
	var result = make(map[int]bool)
	for i, cell := range row.Cells {
		if len(cell.String()) != 0 {
			result[i] = true
		}
	}
	return result
}

func DumpRow(row *xlsx.Row, sheetname string, flags map[int]bool, headers map[int]string, id int) {
	idcell := ""
	for i, cell := range row.Cells {
		if i == id - 1 {
			idcell = cell.String()
			continue
		}
		if len(idcell) == 0 {
			continue
		}
		header, ok := headers[i]
		if !ok {
			continue
		}
		_, ok = flags[i]
		if !ok {
			continue
		}
		var value = cell.String()
		if len(value) != 0 {
			fmt.Printf("%s(%s) %s: %s\n", sheetname, idcell, header, value)
		}
	}
}
