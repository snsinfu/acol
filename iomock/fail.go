package iomock

import (
	"errors"
)

/*
FailingIO implements io.Reader and io.Writer and always fails on everything.
*/
type FailingIO struct {
}

/*
Read just fails with iomock.ErrorMessage.
*/
func (reader *FailingIO) Read(_ []byte) (int, error) {
	return 0, errors.New(ErrorMessage)
}

/*
Write just fails with iomock.ErrorMessage.
*/
func (reader *FailingIO) Write(_ []byte) (int, error) {
	return 0, errors.New(ErrorMessage)
}
