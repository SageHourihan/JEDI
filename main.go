package main

import (
	"flag"
)

func main() {

	v := flag.String("v", "Please provide file path", "Used to validate an EDI file")
	t := flag.String("t", "Please provide file path", "Used to translate an EDI file to JSON")
	flag.Parse()

	//? i feel like there is a better way to do this
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "v" {
			validate(*v)//, f.Name)
		} else if f.Name == "t" {
			translate(*t)//, f.Name)
		}
	})
}


