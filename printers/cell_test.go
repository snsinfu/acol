package printers

import (
	"reflect"
	"testing"
)

func Test_MakeCells(t *testing.T) {
	testCases := []struct {
		items    []string
		expected []Cell
	}{
		{
			items:    []string{},
			expected: []Cell{},
		},
		{
			items:    []string{"abc"},
			expected: []Cell{{"abc", 3}},
		},
		{
			items:    []string{"abc", "defg", "hijkl"},
			expected: []Cell{{"abc", 3}, {"defg", 4}, {"hijkl", 5}},
		},
	}
	for _, testCase := range testCases {
		actual := MakeCells(testCase.items)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v => %v, want %v",
				testCase.items, actual, testCase.expected)
		}
	}
}
