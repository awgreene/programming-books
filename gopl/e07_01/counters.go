package e07_01

import (
	"bufio"
	"bytes"
	"io"
)

type ByteCounter int

type WordCounter int

type LineCounter int

func splitCounter(r io.Reader, split bufio.SplitFunc) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(split)
	var count int
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func (bc *ByteCounter) Write(b []byte) (int, error) {
	n, err := splitCounter(bytes.NewBuffer(b), bufio.ScanBytes)
	*bc += ByteCounter(n)
	return n, err
}

func (wc *WordCounter) Write(b []byte) (int, error) {
	n, err := splitCounter(bytes.NewBuffer(b), bufio.ScanWords)
	*wc += WordCounter(n)
	return n, err
}

func (c *LineCounter) Write(b []byte) (int, error) {
	n, err := splitCounter(bytes.NewBuffer(b), bufio.ScanLines)
	*c += LineCounter(n)
	return n, err
}
