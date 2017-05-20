package printers

import (
	"github.com/mattn/go-runewidth"
)

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
		cells[i].Width = runewidth.StringWidth(item)
	}
	return cells
}
