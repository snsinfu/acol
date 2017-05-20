package printers

import (
	"reflect"
	"testing"
)

func TestDenseColumnMajor_determineColumnWidths(t *testing.T) {
	testCases := []struct {
		width    int
		spacing  int
		cells    []textCell
		expected []int
	}{
		// Degenerate
		{
			width:    0,
			spacing:  0,
			cells:    []textCell{},
			expected: []int{0},
		},
		// Overflow (width == 0)
		{
			width:   0,
			spacing: 0,
			cells: []textCell{
				{"", 10},
			},
			expected: []int{10},
		},
		// Overflow (width > 0)
		{
			width:   9,
			spacing: 0,
			cells: []textCell{
				{"", 10},
			},
			expected: []int{10},
		},
		// Overflow (due to spacing)
		{
			width:   20,
			spacing: 1,
			cells: []textCell{
				{"", 10}, {"", 10},
			},
			expected: []int{10},
		},
		// Single column
		{
			width:   10,
			spacing: 0,
			cells: []textCell{
				{"", 5}, {"", 6}, {"", 7},
			},
			expected: []int{7},
		},
		// Two columns
		{
			width:   15,
			spacing: 1,
			cells: []textCell{
				// [5] [7]
				// [6]
				{"", 5}, {"", 6}, {"", 7},
			},
			expected: []int{6, 7},
		},
		// Three columns
		{
			width:   25,
			spacing: 1,
			cells: []textCell{
				// [5] [7] [9]
				// [6] [8]
				{"", 5}, {"", 6}, {"", 7}, {"", 8}, {"", 9},
			},
			expected: []int{6, 8, 9},
		},
	}
	for _, testCase := range testCases {
		printer := NewDenseColumnMajor(testCase.width, testCase.spacing)
		actual := printer.determineColumnWidths(testCase.cells)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v, %v | %v => %v, want %v",
				testCase.width, testCase.spacing, testCase.cells, actual, testCase.expected)
		}
	}
}
