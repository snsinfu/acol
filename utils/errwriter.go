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
Err returns error occured on previous operations.
*/
func (this *ErrWriter) Err() error {
	return this.err
}

/*
WriteString writes given string. It does nothing if Err is not nil.
*/
func (this *ErrWriter) WriteString(str string) {
	if this.err != nil {
		return
	}
	_, err := this.writer.WriteString(str)
	this.err = err
}

/*
Flush flushes the internal buffer. It does nothing if Err is not nil.
*/
func (this *ErrWriter) Flush() {
	if this.err != nil {
		return
	}
	this.err = this.writer.Flush()
}
