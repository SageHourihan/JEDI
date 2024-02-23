package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/pbjer/hedi"
)

func validate(file string) {

	f, _ := os.ReadFile(file)

	blah :=  string(f)

	reader := strings.NewReader(blah)
	fmt.Println(reader)
	lexer := hedi.NewLexer(reader)
	tokens, err := lexer.Tokens()

	if err != nil {
		fmt.Println(tokens)
	}
}

func translate(file string) {
	fmt.Println("Translate logic")
}
