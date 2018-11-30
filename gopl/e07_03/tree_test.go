package e07_03_test

import (
	"testing"

	"github.com/awgreene/books/gopl/e07_03"
)

func TestTree(t *testing.T) {
	tree := e07_03.Add(nil, 2)
	tree = e07_03.Add(tree, 3)
	tree = e07_03.Add(tree, 1)
	tree = e07_03.Add(tree, 4)

	s := tree.String()
	expected := "[2, 1, 3, 4, ]"
	if s != expected {
		t.Errorf("String() returned %s, expected: %s", s, expected)
	}

}
