package e07_04_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/awgreene/books/gopl/e07_04"
	"golang.org/x/net/html"
)

const sampleHTML = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Test page for parsing</title>
  </head>
  <body>
    <p>Some content here</p>
  </body>
</html>
`

func TestStringReaderNilP(t *testing.T) {
	var b []byte
	reader := e07_04.NewReader("0987654321")

	n, err := reader.Read(b)
	if n != 0 || err != nil {
		t.Errorf("Error while reading to nil slice. n: %d, err: %v, b: %v",
			n, err, b)
	}
}

func TestStringReaderEmptyReaderNilP(t *testing.T) {
	var b []byte
	reader := e07_04.NewReader("")

	n, err := reader.Read(b)
	if n != 0 || err != io.EOF {
		t.Errorf("Error while reading to nil slice empty string. n: %d, err: %v, p: %v",
			n, err, b)
	}

	b = make([]byte, 5)
	n, err = reader.Read(b)
	if n != 0 || err != io.EOF {
		t.Errorf("Error while reading to nil slice empty string. n: %d, err: %v, p: %v",
			n, err, b)
	}
}
func TestStringReaderSmallerDest(t *testing.T) {
	b := make([]byte, 2)

	reader := e07_04.NewReader("0987654321")

	n, err := reader.Read(b)
	if n != 2 || err != nil || bytes.Compare(b, []byte("09")) != 0 {
		t.Errorf("Error while reading. n: %d, err: %v, p: %v",
			n, err, b)
	}
}

func TestStringReaderLargerDest(t *testing.T) {
	b := make([]byte, 11)

	reader := e07_04.NewReader("0987654321")

	n, err := reader.Read(b)
	if n != 10 || err != nil || bytes.Compare(b, []byte("0987654321\x00")) != 0 {
		t.Errorf("Error while reading. n: %d, err: %v, p: %v",
			n, err, b)
	}
}

func TestStringReaderSameSizeSrcAndDest(t *testing.T) {
	b := make([]byte, 5)

	reader := e07_04.NewReader("0987654321")

	n, err := reader.Read(b)
	if n != 5 || err != nil || bytes.Compare(b, []byte("09876")) != 0 {
		t.Errorf("Error while reading. n: %d, err: %v, p: %v",
			n, err, b)
	}

	n, err = reader.Read(b)
	if n != 5 || err != nil || bytes.Compare(b, []byte("54321")) != 0 {
		t.Errorf("Error while reading. n: %d, err: %v, p: %v",
			n, err, b)
	}

	n, err = reader.Read(b)
	if n != 0 || err != io.EOF || bytes.Compare(b, []byte("54321")) != 0 {
		t.Errorf("Error reading after file. n: %d, err: %v, p: %v",
			n, err, b)
	}
}

func TestStringReaderHTML(t *testing.T) {
	_, err := html.Parse(e07_04.NewReader(sampleHTML))
	if err != nil {
		t.Errorf("Failed to parse html from String Reader.")
	}
}
