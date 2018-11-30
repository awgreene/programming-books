package e07_05

import (
	"io"
)

type limitReader struct {
	reader          io.Reader
	limit, position int64
}

// LimitReader returns a new Reader reading from str.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{reader: r, limit: n}
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	var readLen int64

	if lr.position == lr.limit && len(p) > 0 {
		return 0, io.EOF
	}
	/*
		limit = 10
		position 6
		len(p) 5
	*/

	if lr.limit < lr.position+int64(len(p)) {
		readLen = lr.limit - lr.position
	} else {
		readLen = int64(len(p))
	}

	n, err = lr.reader.Read(p[:readLen])
	lr.position += int64(n)

	return n, err
}
