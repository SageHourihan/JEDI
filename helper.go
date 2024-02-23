package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url string = "https://api.edination.com/v2/x12/read?ignoreNullValues=false&continueOnError=false&charSet=utf-8"
const method string = "POST"

func validate(file string) {

	f, _ := os.ReadFile(file)

	payload := strings.NewReader(string(f))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", "3ecf6b1c5cf34bd797a5f4c57951a1cf")
	req.Header.Add("Content-Type", "application/octet-stream")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func translate(file string) {
	fmt.Println("Translate logic")
}
