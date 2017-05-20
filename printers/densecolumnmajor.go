package printers

import (
	"fmt"

	"github.com/frickiericker/acol/utils"
)

/*
DenseColumnMajor is a printer that displays items in a column-major table
format.
*/
type DenseColumnMajor struct {
	outputWidth   int
	columnSpacing int
}

/*
NewDenseColumnMajor creates a DenseColumnMajor printer. The output format will
be a table of given `width` in which columns are separated by `spacing` spaces.
*/
func NewDenseColumnMajor(width, spacing int) *DenseColumnMajor {
	return &DenseColumnMajor{
		outputWidth:   width,
		columnSpacing: spacing,
	}
}

/*
Print displays items in a column-major table format.
*/
func (this *DenseColumnMajor) Print(items []string) {
	cells := makeTextCells(items)
	columnWidths := this.determineColumnWidths(cells)
	this.printColumns(cells, columnWidths)
}

func (this *DenseColumnMajor) printColumns(cells []textCell, columnWidths []int) {
	numColumns := len(columnWidths)
	numRows := (len(cells) + numColumns - 1) / numColumns
	fmt.Println(numColumns, numRows) // FIXME
}

func (this *DenseColumnMajor) determineColumnWidths(cells []textCell) []int {
	maxColumns := utils.IntMax(1, utils.IntMin(this.outputWidth, len(cells)))
	columnWidths := make([]int, maxColumns)
	for numColumns := maxColumns; numColumns > 1; numColumns-- {
		columnWidths = columnWidths[:numColumns]
		if this.computeAndCheckLayout(cells, columnWidths) {
			return columnWidths
		}
	}
	columnWidths = columnWidths[:1]
	this.computeColumnWidths(cells, len(cells), columnWidths)
	return columnWidths
}

func (this *DenseColumnMajor) computeAndCheckLayout(cells []textCell, columnWidths []int) bool {
	// We need to determine the number of rows. This gets rather tricky.
	//
	// Let n be the number of items, c be the number of columns, and r be the
	// number of rows. Then, the number of items in the rightmost column x is
	// given by
	//
	//     (c-1) r + x = n .
	//
	// But x must satisfy 1 <= x <= r for a valid table geometry. This
	// constraint alongside with the equation above gives the possible range
	// for r:
	//
	//     n/c <= r <= (n-1)/(c-1) .
	//
	// We search the number of rows within this range. The smaller the denser,
	// so start with the lower bound.
	numItems := len(cells)
	numColumns := len(columnWidths)
	minRows := (numItems + numColumns - 1) / numColumns
	maxRows := (numItems - 1) / (numColumns - 1)
	for numRows := minRows; numRows <= maxRows; numRows++ {
		this.computeColumnWidths(cells, numRows, columnWidths)
		if this.isValidLayout(columnWidths) {
			return true
		}
	}
	return false
}

func (this *DenseColumnMajor) isValidLayout(columnWidths []int) bool {
	numColumns := len(columnWidths)
	computedWidth := utils.IntSum(columnWidths) + (numColumns-1)*this.columnSpacing
	return computedWidth <= this.outputWidth
}

func (this *DenseColumnMajor) computeColumnWidths(cells []textCell, numRows int, columnWidths []int) {
	for i := range columnWidths {
		columnWidths[i] = 0
	}
	for i, cell := range cells {
		column := i / numRows
		columnWidths[column] = utils.IntMax(columnWidths[column], cell.Width)
	}
}
