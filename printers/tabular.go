package printers

import (
	"github.com/snsinfu/acol/utils"
)

type tabular interface {
	GetOutputWidth() int
	GetColumnSpacing() int
}

type tableShape struct {
	NumColumns   int
	NumRows      int
	ColumnWidths []int
}

func isValidTableShape(printer tabular, shape tableShape) bool {
	columnSpacing := printer.GetColumnSpacing()
	computedWidth := utils.IntSum(shape.ColumnWidths) + (shape.NumColumns-1)*columnSpacing
	return computedWidth <= printer.GetOutputWidth()
}
