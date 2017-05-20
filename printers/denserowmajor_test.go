package printers

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDenseRowMajor_Print(t *testing.T) {
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
			expected: "the   quick\nbrown fox\njumps\n",
		},
		{
			width:    15,
			spacing:  1,
			cells:    []Cell{{"the", 3}, {"quick", 5}, {"brown", 5}, {"fox", 3}, {"jumps", 5}},
			expected: "the quick brown\nfox jumps\n",
		},
	}
	for _, testCase := range testCases {
		printer := NewDenseRowMajor(testCase.width, testCase.spacing)
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

func TestDenseRowMajor_determineColumnWidths(t *testing.T) {
	testCases := []struct {
		width    int
		spacing  int
		cells    []Cell
		expected []int
	}{
		// Degenerate
		{
			width:    0,
			spacing:  0,
			cells:    []Cell{},
			expected: []int{0},
		},
		// Overflow (width == 0)
		{
			width:   0,
			spacing: 0,
			cells: []Cell{
				{"", 10},
			},
			expected: []int{10},
		},
		// Overflow (width > 0)
		{
			width:   9,
			spacing: 0,
			cells: []Cell{
				{"", 10},
			},
			expected: []int{10},
		},
		// Overflow (due to spacing)
		{
			width:   20,
			spacing: 1,
			cells: []Cell{
				{"", 10},
				{"", 10},
			},
			expected: []int{10},
		},
		// Okay
		{
			width:   30,
			spacing: 1,
			cells: []Cell{
				{"", 5}, {"", 6}, {"", 7},
			},
			expected: []int{5, 6, 7},
		},
		// Extension from second row
		{
			width:   30,
			spacing: 1,
			cells: []Cell{
				{"", 10}, {"", 12},
				{"", 11},
			},
			expected: []int{11, 12},
		},
		// Extension from second and third row
		{
			width:   30,
			spacing: 1,
			cells: []Cell{
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
			cells: []Cell{
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
