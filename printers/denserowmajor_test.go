package printers

import (
	"reflect"
	"testing"
)

func TestDenseRowMajor_determineColumnWidths(t *testing.T) {
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
				{"", 10},
				{"", 10},
			},
			expected: []int{10},
		},
		// Okay
		{
			width:   30,
			spacing: 1,
			cells: []textCell{
				{"", 5}, {"", 6}, {"", 7},
			},
			expected: []int{5, 6, 7},
		},
		// Extension from second row
		{
			width:   30,
			spacing: 1,
			cells: []textCell{
				{"", 10}, {"", 12},
				{"", 11},
			},
			expected: []int{11, 12},
		},
		// Extension from second and third row
		{
			width:   30,
			spacing: 1,
			cells: []textCell{
				{"", 10}, {"", 12},
				{"", 11}, {"", 11},
				{"", 10}, {"", 13},
			},
			expected: []int{11, 13},
		},
		// No extension
		{
			width:   30,
			spacing: 1,
			cells: []textCell{
				{"", 13}, {"", 15},
				{"", 11}, {"", 11},
				{"", 10}, {"", 13},
			},
			expected: []int{13, 15},
		},
	}
	for _, testCase := range testCases {
		printer := NewDenseRowMajor(testCase.width, testCase.spacing)
		actual := printer.determineColumnWidths(testCase.cells)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v, %v | %v => %v, want %v",
				testCase.width, testCase.spacing, testCase.cells, actual, testCase.expected)
		}
	}
}
