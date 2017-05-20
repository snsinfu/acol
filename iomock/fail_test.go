package iomock

import (
	"testing"
)

func TestFailingIO_Read(t *testing.T) {
	io := new(FailingIO)
	data := make([]byte, 10)
	_, err := io.Read(data)
	if err == nil {
		t.Error("unexpected success")
	}
	if err != nil && err != Error {
		t.Error("unexpected error:", err)
	}
}

func TestFailingIO_Write(t *testing.T) {
	io := new(FailingIO)
	data := []byte{1, 2, 3}
	_, err := io.Write(data)
	if err == nil {
		t.Error("unexpected success")
	}
	if err != nil && err != Error {
		t.Error("unexpected error:", err)
	}
}
