package printers

import (
	"github.com/mattn/go-runewidth"
)

type textCell struct {
	Content string
	Width   int
}

func makeTextCells(items []string) []textCell {
	cells := make([]textCell, len(items))
	for i, item := range items {
		cells[i].Content = item
		cells[i].Width = runewidth.StringWidth(item)
	}
	return cells
}
