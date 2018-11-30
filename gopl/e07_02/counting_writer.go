package e07_02

import (
	"io"
)

type writerCounter struct {
	w io.Writer
	c int64
}

func (wc *writerCounter) Write(p []byte) (n int, err error) {
	wc.c += int64(len(p))
	return wc.w.Write(p)
}

/*
 A function that, given an io.Writer, returns a new Writer that wraps
 the original, and a pointer to an int64 variable that at any moment
 contains the number of bytes writtern to the new Writer.
*/
func CountingWriter(cw io.Writer) (io.Writer, *int64) {
	w := writerCounter{cw, 0}
	return &w, &(w.c)
}
