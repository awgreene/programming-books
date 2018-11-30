package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type fn func()

// Exercise 1.3
// Experiment to measure the difference in running time between our potentially inefficent versions and the one that uses `strings.join`
func main() {
	test(echo1, 1)
	test(echo2, 2)
	test(echo3, 3)
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := " ", " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func runtimeReport(test int, duration float64) {
	fmt.Printf("Echo%v took %v seconds\n", test, duration)
}

func test(f fn, i int) {
	t := time.Now()
	f()
	runtimeReport(i, time.Since(t).Seconds())
}
