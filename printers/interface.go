package printers

import (
	"io"
)

/*
Common interface of printers. A printer writes given items (slice of strings) to
a Writer in a specific format.
*/
type Interface interface {
	Print(out io.Writer, items []string)
}
