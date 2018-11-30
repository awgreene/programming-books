package e07_05_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/awgreene/books/gopl/e07_05"
)

func TestReader(t *testing.T) {
	r := e07_05.LimitReader(strings.NewReader("ABCDE"), 3)
	p := make([]byte, 2)

	n, err := r.Read(p)
	if n != 2 || err != nil || bytes.Compare(p, []byte("AB")) != 0 {
		t.Errorf("Unexpected read size")
	}

	total := n
	n, err = r.Read(p)
	total += n
	if total != 3 || err != nil || bytes.Compare(p, []byte("CB")) != 0 {
		t.Errorf("Off limits read. total: %d, err: %v, p: %v", total, err, p)
	}

	n, err = r.Read(p)
	if n != 0 || err != io.EOF || bytes.Compare(p, []byte("CB")) != 0 {
		t.Errorf("Off limits read")
	}
}
