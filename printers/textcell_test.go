package printers

import (
	"reflect"
	"testing"
)

func Test_makeTextCells(t *testing.T) {
	testCases := []struct {
		items    []string
		expected []textCell
	}{
		{
			items:    []string{},
			expected: []textCell{},
		},
		{
			items:    []string{"abc"},
			expected: []textCell{{"abc", 3}},
		},
		{
			items:    []string{"abc", "defg", "hijkl"},
			expected: []textCell{{"abc", 3}, {"defg", 4}, {"hijkl", 5}},
		},
	}
	for _, testCase := range testCases {
		actual := makeTextCells(testCase.items)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v => %v, want %v",
				testCase.items, actual, testCase.expected)
		}
	}
}
