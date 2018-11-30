package e07_02_test

import (
	"bytes"
	"testing"

	"github.com/awgreene/books/gopl/e07_02"
)

func TestWriterCounter(t *testing.T) {
	buf := new(bytes.Buffer)
	w, cPtr := e07_02.CountingWriter(buf)

	s := "54321"
	w.Write([]byte(s))
	if *cPtr != 5 {
		t.Errorf("Number of bytes counted should be 10: %d", *cPtr)
	}
	if buf.String() != "54321" {
		t.Errorf("Buffer shoud be equals to '%s'", s)
	}

	w.Write([]byte(""))
	if *cPtr != 5 {
		t.Errorf("Number of bytes counted should be 5: %d", *cPtr)
	}

	w.Write([]byte("123"))
	if *cPtr != 8 {
		t.Errorf("Number of bytes counted should be 8: %d", *cPtr)
	}

	w.Write([]byte(nil))
	if *cPtr != 8 {
		t.Errorf("Number of bytes counted should be 8: %d", *cPtr)
	}
}
