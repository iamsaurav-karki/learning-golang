package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Making GET Request")
	performGetRequest()
}

func performGetRequest() {
	const url = "https://sauravverse.xyz"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response status code:", response.StatusCode)
	fmt.Println("Response content length:", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Response string:", responseString.String())

	// fmt.Println("Response body:", string(content))
}
