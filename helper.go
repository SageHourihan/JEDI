package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// struct to store our json response in
type Detail struct {
	Index            int         `json:"Index"`
	TransactionIndex int         `json:"TransactionIndex"`
	TransactionRef   string      `json:"TransactionRef"`
	SpecRef          string      `json:"SpecRef"`
	SegmentId        string      `json:"SegmentId"`
	DataElementId    interface{} `json:"DataElementId"`
	DataElementIndex int         `json:"DataElementIndex"`
	CompositeIndex   int         `json:"CompositeIndex"`
	RepIndex         int         `json:"RepIndex"`
	Value            interface{} `json:"Value"`
	Message          string      `json:"Message"`
	Status           string      `json:"Status"`
}

type Response struct {
	LastIndex int      `json:"LastIndex"`
	Details   []Detail `json:"Details"`
	Status    string   `json:"Status"`
}

const x12_read string = "https://api.edination.com/v2/x12/read?ignoreNullValues=false&continueOnError=false&charSet=utf-8"
const x12_validate string = "https://api.edination.com/v2/x12/validate?basicSyntax=false&skipTrailer=false&structureOnly=false"
const method string = "POST"

func read(file string) {

	f, _ := os.ReadFile(file)

	payload := strings.NewReader(string(f))

	client := &http.Client{}
	req, err := http.NewRequest(method, x12_read, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", "3ecf6b1c5cf34bd797a5f4c57951a1cf")
	req.Header.Add("Content-Type", "application/octet-stream")

	// Send the request to the API
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read and print the API response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert response from read to
	edi := string(body)
	//TODO: add validate function
	// rename read to translate possibly. 

}

func translate(file string) {
	fmt.Println("Translate logic")
}
