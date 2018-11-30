package e07_04

import (
	"io"
)

type stringReader struct {
	position int
	str      string
}

// NewReader returns a new Reader reading from str.
func NewReader(str string) io.Reader {
	return &stringReader{str: str, position: 0}
}

func (sr *stringReader) Read(b []byte) (n int, err error) {
	n = copy(b, sr.str[sr.position:])
	sr.position += n

	if n == 0 && sr.position == len(sr.str) {
		err = io.EOF
	}

	return n, err
}
