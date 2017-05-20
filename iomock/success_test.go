package iomock

import (
	"testing"
)

func TestSucceedingIO_Read(t *testing.T) {
	io := new(SucceedingIO)
	data := make([]byte, 10)
	_, err := io.Read(data)
	if err != nil {
		t.Error("unexpected error:", err)
	}
}

func TestSucceedingIO_Write(t *testing.T) {
	io := new(SucceedingIO)
	data := []byte{1, 2, 3}
	_, err := io.Write(data)
	if err != nil {
		t.Error("unexpected error:", err)
	}
}
