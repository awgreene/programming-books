package main

import (
	"flag"
	"fmt"
)

var temp = CelsiusFlag("temp", 100.0, "the Temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
