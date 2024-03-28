package main

import (
	// "fmt"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/arcward/edx12"
)

func translate(messageText string) {

	outputFile, err := os.Create("./tmp/output.txt")

	if err != nil {
		log.Panicln(err)
	}

	dat, _ := os.ReadFile(messageText)

	rawMessage, _ := edx12.Read([]byte(dat))
	message, _ := rawMessage.Message(context.Background())

	encoder := json.NewEncoder(outputFile)
	encoder.SetEscapeHTML(false) // avoid '>' being escaped to '\u003e'
	encoder.SetIndent("", "  ")
	err = encoder.Encode(message)

	if err != nil {
		log.Panicln(err)
	}
}


func validate(messageText string ) {

	dat, _ := os.ReadFile(messageText)

	rawMessage := edx12.Validate([]byte(dat))

	fmt.Println(rawMessage)
}