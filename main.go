package main

import (
	"flag"
)

func main() {

	v := flag.String("v", "Please provide file path", "Used to validate an EDI file")
	t := flag.String("t", "Please provide file path", "Used to translate an EDI file")
	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "v" {
			read(*v)
		} else if f.Name == "t" {
			translate(*t)
		}
	})
}


