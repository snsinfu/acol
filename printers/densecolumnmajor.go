package printers

import (
	"io"
	"strings"

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
GetOutputWidth returns the width of the table.
*/
func (this *DenseColumnMajor) GetOutputWidth() int {
	return this.outputWidth
}

/*
GetColumnSpacing returns the space between columns.
*/
func (this *DenseColumnMajor) GetColumnSpacing() int {
	return this.columnSpacing
}

/*
Print displays cells in a column-major table format.
*/
func (this *DenseColumnMajor) Print(out io.Writer, cells []Cell) {
	shape := this.determineShape(cells)
	maxSpacing := utils.IntMaxReduce(shape.ColumnWidths, 0) + this.columnSpacing
	cachedSpaces := strings.Repeat(" ", maxSpacing)

	writer := utils.NewErrWriter(out)
	// We need to print cells in row-major order while the cells are stored in
	// column-major order.
	for row := 0; row < shape.NumRows; row++ {
		for column := 0; column < shape.NumColumns; column++ {
			i := column*shape.NumRows + row
			if i >= len(cells) {
				break
			}
			cell := cells[i]
			writer.WriteString(cell.Content)
			if i+shape.NumRows < len(cells) {
				padding := shape.ColumnWidths[column] - cell.Width
				spacing := padding + this.columnSpacing
				writer.WriteString(cachedSpaces[:spacing])
			} else {
				writer.WriteString("\n")
			}
		}
	}
	writer.Flush()
}

func (this *DenseColumnMajor) determineShape(cells []Cell) tableShape {
	maxColumns := utils.IntMax(1, utils.IntMin(this.outputWidth, len(cells)))
	columnWidths := make([]int, maxColumns)
	for numColumns := maxColumns; numColumns > 1; numColumns-- {
		shape := tableShape{
			ColumnWidths: columnWidths[:numColumns],
		}
		if this.tryComputeShape(cells, numColumns, &shape) {
			return shape
		}
	}
	return this.getFallbackShape(cells)
}

func (this *DenseColumnMajor) getFallbackShape(cells []Cell) tableShape {
	shape := tableShape{
		NumColumns:   1,
		NumRows:      len(cells),
		ColumnWidths: make([]int, 1),
	}
	this.computeColumnWidths(cells, shape.NumRows, shape.ColumnWidths)
	return shape
}

func (this *DenseColumnMajor) tryComputeShape(cells []Cell, numColumns int, shape *tableShape) bool {
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
	minRows := (numItems + numColumns - 1) / numColumns
	maxRows := (numItems - 1) / (numColumns - 1)
	for numRows := minRows; numRows <= maxRows; numRows++ {
		shape.NumColumns = numColumns
		shape.NumRows = numRows
		this.computeColumnWidths(cells, numRows, shape.ColumnWidths)
		if isValidTableShape(this, *shape) {
			return true
		}
	}
	return false
}

func (this *DenseColumnMajor) computeColumnWidths(cells []Cell, numRows int, widths []int) {
	for i := range widths {
		widths[i] = 0
	}
	for i, cell := range cells {
		column := i / numRows
		widths[column] = utils.IntMax(widths[column], cell.Width)
	}
}
