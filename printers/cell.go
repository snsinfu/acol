package printers

import (
	"github.com/snsinfu/acol/descape"
)

/*
Cell is text with width information.
*/
type Cell struct {
	Content string
	Width   int
}

/*
MakeCells maps slice of strings into cells using go-runewidth package.
*/
func MakeCells(items []string) []Cell {
	cells := make([]Cell, len(items))
	for i, item := range items {
		cells[i].Content = item
		cells[i].Width = descape.StringWidth(item)
	}
	return cells
}
