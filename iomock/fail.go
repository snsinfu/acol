package iomock

import "errors"

/*
Error is used as mocked error.
*/
var Error = errors.New("this is a false error from a mock")

/*
FailingIO implements io.Reader and io.Writer and always fails on everything.
*/
type FailingIO struct {
}

/*
Read just fails with iomock.ErrorMessage.
*/
func (reader *FailingIO) Read(_ []byte) (int, error) {
	return 0, Error
}

/*
Write just fails with iomock.ErrorMessage.
*/
func (reader *FailingIO) Write(_ []byte) (int, error) {
	return 0, Error
}
