package functors

import (
	"io"
)

type WriterMock struct {
	values []string
}

func NewWriterMock() *WriterMock {
	return &WriterMock{
		values: make([]string, 0),
	}
}

func (w *WriterMock) Write(newValue []byte) (n int, err error) {
	w.values = append(w.values, string(newValue))
	return
}

func (w *WriterMock) GetValues() []string {
	return w.values
}

func OutputWrite(input io.Writer, message string) {
	input.Write([]byte(message))
}
