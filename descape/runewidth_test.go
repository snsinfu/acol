package descape

import (
	"testing"
)

func TestStringWidth(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		// No escape sequence
		{"", 0},
		{"a", 1},
		{"文字", 4},
		{"文letter字", 10},
		// Affixed by escape sequences
		{"\x1b[45;33;1m\x1b[m", 0},
		{"\x1b[45;33;1mABC\x1b[m", 3},
		{"\x1b[45;33;1m文字\x1b[m", 4},
		// Escape sequences at the middle of text
		{"The quick brown \x1b[45;33;1mfox\x1b[m jumps over", 30},
	}
	for _, testCase := range testCases {
		actual := StringWidth(testCase.input)
		if actual != testCase.expected {
			t.Errorf("%v => %v, want %v", testCase.input, actual, testCase.expected)
		}
	}
}
