package utils

import (
	"bufio"
	"io"
)

/*
ErrWriter does buffered output with deferred error check.
*/
type ErrWriter struct {
	writer *bufio.Writer
	err    error
}

/*
NewErrWriter creates a new ErrWriter that writes to given Writer.
*/
func NewErrWriter(writer io.Writer) *ErrWriter {
	return &ErrWriter{
		writer: bufio.NewWriter(writer),
		err:    nil,
	}
}

/*
Err returns error occurred on previous operations.
*/
func (ewriter *ErrWriter) Err() error {
	return ewriter.err
}

/*
WriteString writes given string. It does nothing if Err is not nil.
*/
func (ewriter *ErrWriter) WriteString(str string) {
	if ewriter.err != nil {
		return
	}
	_, err := ewriter.writer.WriteString(str)
	ewriter.err = err
}

/*
Flush flushes the internal buffer. It does nothing if Err is not nil.
*/
func (ewriter *ErrWriter) Flush() {
	if ewriter.err != nil {
		return
	}
	ewriter.err = ewriter.writer.Flush()
}
