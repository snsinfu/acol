package printers

import (
	"io"
)

/*
Common interface of printers. A printer writes given cells (strings with width
information) to a Writer in a specific format.
*/
type Interface interface {
	Print(out io.Writer, cells []Cell)
}
