package utils

import (
	"errors"
	"testing"
)

type succeedingWriter struct {
}

func (this *succeedingWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

type failingWriter struct {
}

func (this *failingWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("this is a test")
}

func TestErrWriter_withoutError(t *testing.T) {
	writer := NewErrWriter(new(succeedingWriter))
	if err := writer.Err(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	writer.WriteString("the quick brown fox jumps over the lazy dog")
	writer.Flush()
	if err := writer.Err(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestErrWriter_withError(t *testing.T) {
	writer := NewErrWriter(new(failingWriter))
	if err := writer.Err(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	writer.WriteString("the quick brown fox jumps over the lazy dog")
	writer.Flush()
	if err := writer.Err(); err == nil {
		t.Error("error is expected")
	}
	if err := writer.Err(); err.Error() != "this is a test" {
		t.Errorf("unexpected error: %v", err)
	}
}
