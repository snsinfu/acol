package printers

import (
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
GetOutputWidth returns the width of the table.
*/
func (this *DenseRowMajor) GetOutputWidth() int {
	return this.outputWidth
}

/*
GetColumnSpacing returns the space between columns.
*/
func (this *DenseRowMajor) GetColumnSpacing() int {
	return this.columnSpacing
}

/*
Print displays cells in a row-major table format.
*/
func (this *DenseRowMajor) Print(out io.Writer, cells []Cell) {
	shape := this.determineShape(cells)
	maxSpacing := utils.IntMaxReduce(shape.ColumnWidths, 0) + this.columnSpacing
	cachedSpaces := strings.Repeat(" ", maxSpacing)

	writer := utils.NewErrWriter(out)
	for i, cell := range cells {
		writer.WriteString(cell.Content)
		column := i % shape.NumColumns
		if column == shape.NumColumns-1 || i == len(cells)-1 {
			writer.WriteString("\n")
		} else {
			padding := shape.ColumnWidths[column] - cells[i].Width
			spacing := padding + this.columnSpacing
			writer.WriteString(cachedSpaces[:spacing])
		}
	}
	writer.Flush()
}

func (this *DenseRowMajor) determineShape(cells []Cell) tableShape {
	maxColumns := utils.IntMax(1, utils.IntMin(this.outputWidth/2, len(cells)))
	columnWidths := make([]int, maxColumns)
	for numColumns := maxColumns; numColumns > 0; numColumns-- {
		shape := tableShape{
			NumRows:      (len(cells) + numColumns - 1) / numColumns,
			NumColumns:   numColumns,
			ColumnWidths: columnWidths[:numColumns],
		}
		this.computeColumnWidths(cells, shape.ColumnWidths)
		if isValidTableShape(this, shape) {
			return shape
		}
	}
	return this.getFallbackShape(cells)
}

func (this *DenseRowMajor) getFallbackShape(cells []Cell) tableShape {
	shape := tableShape{
		NumColumns:   1,
		NumRows:      len(cells),
		ColumnWidths: make([]int, 1),
	}
	this.computeColumnWidths(cells, shape.ColumnWidths)
	return shape
}

func (this *DenseRowMajor) computeColumnWidths(cells []Cell, widths []int) {
	for i := range widths {
		widths[i] = 0
	}
	for i, cell := range cells {
		column := i % len(widths)
		widths[column] = utils.IntMax(widths[column], cell.Width)
	}
}
