package helpers

import "fmt"

func FormatCell(sheet, symbol string, firstRow int) string {
	return fmt.Sprintf("'%s'!$%s$%d", sheet, symbol, firstRow)
}

func FormatDataRange(sheet, symbolFrom, symbolTo string, firstRow, lastRow int) string {
	return fmt.Sprintf("'%s'!$%s$%d:$%s$%d", sheet, symbolFrom, firstRow+1, symbolTo, lastRow)
}
