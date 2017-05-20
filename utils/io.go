package utils

import (
	"bufio"
	"io"
)

/*
ReadLines reads all lines from given reader as an array. Newline characters are
removed.
*/
func ReadLines(src io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(src)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
