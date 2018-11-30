package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g C°", c) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "K":
		fmt.Println(unit)
		f.Celsius = Celsius(value)
		return nil
	}

	return fmt.Errorf("Invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
