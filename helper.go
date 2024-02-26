package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func translate(file string, mode string) {

	// reading the file
	f, _ := os.ReadFile(file)

	// sending request 
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

	// if mode is translate, translate and write output
	if mode == "t" {

		// declaring var with json
		newJSON := body

		// removing result object from json
		result := gjson.GetBytes(body, "#.Result")
		for i := range result.Array() {
			newJSON, err = sjson.DeleteBytes(newJSON, strconv.Itoa(i)+".Result")
			check(err)
		}

		// creating path to write
		f, err := os.Create("./tmp/output.txt")
		check(err)

		defer f.Close()

		// writing file
		writeString := strings.Trim(string(newJSON), "[]")
		f.WriteString(writeString)

		fmt.Print("file written")
	} else if mode == "v" {
		validate(edi)
	}
}

func validate(edi string) {

	// trim json so its valid for ediNation
	p := strings.Trim(edi, "[]")

	// create a payload with the api respnse
	payload := strings.NewReader(p)

	// create HTTP client
	client := &http.Client{}

	// create HTTP request
	req, err := http.NewRequest(method, x12_validate, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set headers for the API request
	req.Header.Add("Ocp-Apim-Subscription-Key", "3ecf6b1c5cf34bd797a5f4c57951a1cf")
	req.Header.Add("Content-Type", "application/json")

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

	// explicitly declaring response variable
	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if response.Status == "success" {
		fmt.Println("Valid EDI")
	} else if response.Status == "error" {
		message := response.Details[0].Message
		segmentId := response.Details[0].SegmentId
		msg := fmt.Sprintf("Error: %s - %s", message, segmentId)
		fmt.Println(msg)
	}
}
