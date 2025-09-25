package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {

	fmt.Println("Handling url")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{

		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=saurav",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
