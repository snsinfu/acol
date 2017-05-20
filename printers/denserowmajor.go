package printers

import (
	"fmt"
	"io"
	"strings"

	"github.com/frickiericker/acol/utils"
)

/*
DenseRowMajor is a printer that displays items in a row-major table format.
*/
type DenseRowMajor struct {
	outputWidth   int
	columnSpacing int
}

/*
NewDenseRowMajor creates a DenseRowMajor printer. The output format will be a
table of given `width` in which columns are separated by `spacing` spaces.
*/
func NewDenseRowMajor(width, spacing int) *DenseRowMajor {
	return &DenseRowMajor{
		outputWidth:   width,
		columnSpacing: spacing,
	}
}

/*
Print displays items in a row-major table format.
*/
func (this *DenseRowMajor) Print(out io.Writer, items []string) {
	cells := makeTextCells(items)
	columnWidths := this.determineColumnWidths(cells)
	this.printColumns(out, cells, columnWidths)
}

func (this *DenseRowMajor) printColumns(out io.Writer, cells []textCell, columnWidths []int) {
	maxSpacing := utils.IntMaxReduce(columnWidths, 0) + this.columnSpacing
	spaces := strings.Repeat(" ", maxSpacing)
	numColumns := len(columnWidths)
	for i, cell := range cells {
		column := i % numColumns
		if column > 0 {
			padding := columnWidths[column-1] - cells[i-1].Width
			spacing := padding + this.columnSpacing
			fmt.Print(spaces[:spacing])
		}
		fmt.Fprint(out, cell.Content)
		if column == numColumns-1 {
			fmt.Fprint(out, "\n")
		}
	}
	if len(cells)%numColumns != 0 {
		fmt.Fprint(out, "\n")
	}
}

func (this *DenseRowMajor) determineColumnWidths(cells []textCell) []int {
	maxColumns := utils.IntMax(1, utils.IntMin(this.outputWidth/2, len(cells)))
	columnWidths := make([]int, maxColumns)
	for numColumns := maxColumns; numColumns > 0; numColumns-- {
		columnWidths = columnWidths[:numColumns]
		this.computeColumnWidths(cells, columnWidths)
		computedWidth := utils.IntSum(columnWidths) + (numColumns-1)*this.columnSpacing
		if computedWidth <= this.outputWidth {
			break
		}
	}
	return columnWidths
}

func (this *DenseRowMajor) computeColumnWidths(cells []textCell, columnWidths []int) {
	for i := range columnWidths {
		columnWidths[i] = 0
	}
	for i, cell := range cells {
		numColumns := len(columnWidths)
		column := i % numColumns
		columnWidths[column] = utils.IntMax(columnWidths[column], cell.Width)
	}
}
