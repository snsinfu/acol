package printers

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/snsinfu/acol/iomock"
)

func TestDenseColumnMajor_Print(t *testing.T) {
	testCases := []struct {
		width    int
		spacing  int
		cells    []Cell
		expected string
	}{
		{
			width:    0,
			spacing:  0,
			cells:    []Cell{},
			expected: "",
		},
		{
			width:    11,
			spacing:  1,
			cells:    []Cell{{"the", 3}, {"quick", 5}, {"brown", 5}, {"fox", 3}, {"jumps", 5}},
			expected: "the   fox\nquick jumps\nbrown\n",
		},
		{
			width:    17,
			spacing:  1,
			cells:    []Cell{{"the", 3}, {"quick", 5}, {"brown", 5}, {"fox", 3}, {"jumps", 5}},
			expected: "the   brown jumps\nquick fox\n",
		},
	}
	for _, testCase := range testCases {
		printer := NewDenseColumnMajor(testCase.width, testCase.spacing)
		buffer := new(bytes.Buffer)
		printer.Print(buffer, testCase.cells)
		actual := buffer.String()
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v, %v | %v => '%v', want '%v'",
				testCase.width, testCase.spacing, testCase.cells, actual, testCase.expected)
		}
	}
}

func TestDenseColumnMajor_determineShape(t *testing.T) {
	testCases := []struct {
		width    int
		spacing  int
		cells    []Cell
		expected tableShape
	}{
		// Degenerate
		{
			width:   0,
			spacing: 0,
			cells:   []Cell{},
			expected: tableShape{
				NumRows:      0,
				NumColumns:   1,
				ColumnWidths: []int{0},
			},
		},
		// Overflow (width == 0)
		{
			width:   0,
			spacing: 0,
			cells: []Cell{
				{"", 10},
			},
			expected: tableShape{
				NumRows:      1,
				NumColumns:   1,
				ColumnWidths: []int{10},
			},
		},
		// Overflow (width > 0)
		{
			width:   9,
			spacing: 0,
			cells: []Cell{
				{"", 10},
			},
			expected: tableShape{
				NumRows:      1,
				NumColumns:   1,
				ColumnWidths: []int{10},
			},
		},
		// Overflow (due to spacing)
		{
			width:   20,
			spacing: 1,
			cells: []Cell{
				{"", 10}, {"", 10},
			},
			expected: tableShape{
				NumRows:      2,
				NumColumns:   1,
				ColumnWidths: []int{10},
			},
		},
		// Single column
		{
			width:   10,
			spacing: 0,
			cells: []Cell{
				{"", 5}, {"", 6}, {"", 7},
			},
			expected: tableShape{
				NumRows:      3,
				NumColumns:   1,
				ColumnWidths: []int{7},
			},
		},
		// Two columns
		{
			width:   15,
			spacing: 1,
			cells: []Cell{
				// [5] [7]
				// [6]
				{"", 5}, {"", 6}, {"", 7},
			},
			expected: tableShape{
				NumRows:      2,
				NumColumns:   2,
				ColumnWidths: []int{6, 7},
			},
		},
		// Three columns
		{
			width:   25,
			spacing: 1,
			cells: []Cell{
				// [5] [7] [9]
				// [6] [8]
				{"", 5}, {"", 6}, {"", 7}, {"", 8}, {"", 9},
			},
			expected: tableShape{
				NumRows:      2,
				NumColumns:   3,
				ColumnWidths: []int{6, 8, 9},
			},
		},
	}
	for _, testCase := range testCases {
		printer := NewDenseColumnMajor(testCase.width, testCase.spacing)
		actual := printer.determineShape(testCase.cells)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v, %v | %v => %v, want %v",
				testCase.width, testCase.spacing, testCase.cells, actual, testCase.expected)
		}
	}
}

func TestDenseColumnMajor_Print_propagateError(t *testing.T) {
	printer := NewDenseColumnMajor(10, 1)
	writer := new(iomock.FailingIO)
	cells := []Cell{{"abc", 3}, {"def", 3}}
	err := printer.Print(writer, cells)
	if err == nil {
		t.Error("unexpected success")
	}
	if err != nil && err != iomock.Error {
		t.Error("unexpected error:", err)
	}
}
