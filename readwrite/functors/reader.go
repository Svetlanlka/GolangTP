package functors

import "io"

type ReaderMock struct {
	values []string
}

func NewReaderMock() *ReaderMock {
	return &ReaderMock{
		values: make([]string, 0),
	}
}

func (w *ReaderMock) Read(newValue []byte) (n int, err error) {
	w.values = append(w.values, string(newValue))
	return
}

func (w *ReaderMock) GetValues() []string {
	return w.values
}

func OutputRead(output io.Reader, buf string) (n int, err error) {
	return output.Read([]byte(buf))
}
