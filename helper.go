package main

import (
	// "fmt"
	"os"
	"context"
	"encoding/json"

	"github.com/arcward/edx12"
)



func translate(messageText string) {

	dat, _ := os.ReadFile(messageText)

	rawMessage, _ := edx12.Read([]byte(dat))
	message, _ := rawMessage.Message(context.Background())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false) // avoid '>' being escaped to '\u003e'
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(message)
}
