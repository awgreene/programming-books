package main

import (
	"fmt"
	"os"
)

// Exercise 1.2
// Modify the echo program to print the index and value of each of its arguments, one per line.
func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Argument at index", i, "is", os.Args[i])
	}
}
