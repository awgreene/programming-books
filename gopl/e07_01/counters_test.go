package e07_01_test

import (
	"fmt"
	"testing"

	"github.com/awgreene/books/gopl/e07_01"
)

func TestByteCounter(t *testing.T) {

	bc := new(e07_01.ByteCounter)

	n, err := fmt.Fprintf(bc, "Hello there\n General Kenobi")

	if int(*bc) != 27 || err != nil {
		t.Errorf("Failled to count: %d", n)
	}
}

func TestWordCounter(t *testing.T) {

	wc := new(e07_01.WordCounter)

	n, err := fmt.Fprintf(wc, "Hello there\n General Kenobi")

	if int(*wc) != 4 || err != nil {
		t.Errorf("Failled to count: %d", n)
	}
}

func TestLineCounter(t *testing.T) {

	lc := new(e07_01.LineCounter)

	n, err := fmt.Fprintf(lc, "Hello there\n General Kenobi")

	if int(*lc) != 2 || err != nil {
		t.Errorf("Failled to count: %d", n)
	}
}
