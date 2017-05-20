package iomock

/*
SucceedingIO implements io.Reader and io.Writer and always succeeds on
everything.
*/
type SucceedingIO struct {
}

/*
Read just succeeds and returns len(buf) with no error.
*/
func (reader *SucceedingIO) Read(buf []byte) (int, error) {
	return len(buf), nil
}

/*
Write just succeeds and returns len(buf) with no error.
*/
func (reader *SucceedingIO) Write(buf []byte) (int, error) {
	return len(buf), nil
}
