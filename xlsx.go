package exporter

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	defaultColWidth   = 30
	defaultHorizontal = "center"
)

type (
	xlsx struct {
		file  *excelize.File
		sheet string
	}
)

func NewXLSX(numCols int, sty Style) (Generator, error) {
	if numCols <= 0 {
		return nil, fmt.Errorf("number of columns must be positive")
	}

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	x := xlsx{
		file:  f,
		sheet: sheet,
	}

	getColWidth := func() int {
		if sty.ColumnWidth != 0 {
			return sty.ColumnWidth
		}

		return defaultColWidth
	}

	getHorizontal := func() string {
		if sty.Horizontal != "" {
			return sty.Horizontal
		}

		return defaultHorizontal
	}

	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: getHorizontal(),
			WrapText:   true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create center alignment style: %v", err)
	}

	lastCol := x.getColumnName(numCols - 1)
	colRange := fmt.Sprintf("A:%s", lastCol)
	if err := f.SetColStyle(sheet, colRange, style); err != nil {
		return nil, fmt.Errorf("failed to set column style: %v", err)
	}

	if err := f.SetColWidth(sheet, "A", lastCol, float64(getColWidth())); err != nil {
		return nil, fmt.Errorf("failed to set column width: %v", err)
	}

	return &x, nil
}

func (x *xlsx) SetHeaders(headers []string) error {
	for idx, header := range headers {
		cell := fmt.Sprintf("%s1", x.getColumnName(idx))
		if err := x.file.SetCellValue(x.sheet, cell, header); err != nil {
			return fmt.Errorf("failed to set %s header: %v", header, err)
		}
	}

	return nil
}

// getColumnName converts a zero-based column index to Excel column name, e.g., 0 -> A, 27 -> AB
func (x *xlsx) getColumnName(n int) string {
	var col string
	for n >= 0 {
		col = string(rune('A'+(n%26))) + col
		n = n/26 - 1
	}

	return col
}

func (x *xlsx) Generate(entries [][]string) ([]byte, error) {
	for rowIdx, entry := range entries {
		rowNum := rowIdx + 2
		for colIdx, field := range entry {
			if field != "" {
				cell := fmt.Sprintf("%s%d", x.getColumnName(colIdx), rowNum)
				if err := x.file.SetCellValue(x.sheet, cell, field); err != nil {
					return nil, fmt.Errorf("failed to set field: %v", err)
				}
			}
		}
	}

	b, err := x.file.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("failed to write to buffer: %v", err)
	}

	return b.Bytes(), nil
}

func (x *xlsx) Close() error {
	return x.file.Close()
}
