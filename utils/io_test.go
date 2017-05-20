package utils

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestReadLines_examples(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "\n",
			expected: []string{""},
		},
		{
			input:    "abc",
			expected: []string{"abc"},
		},
		{
			input:    "abc\ndefgh",
			expected: []string{"abc", "defgh"},
		},
		{
			input:    "abc\ndefgh\n",
			expected: []string{"abc", "defgh"},
		},
	}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		actual, err := ReadLines(reader)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(
				"%v => %v, want %v",
				testCase.input, actual, testCase.expected)
		}
	}
}

type errorReader struct {
}

func (reader *errorReader) Read(_ []byte) (int, error) {
	return 0, errors.New("expected")
}

func TestReadLines_propagateError(t *testing.T) {
	reader := &errorReader{}
	_, err := ReadLines(reader)
	if err == nil {
		t.Errorf("error not propagated")
	}
	if err.Error() != "expected" {
		t.Errorf("unexpected error: %v", err)
	}
}
