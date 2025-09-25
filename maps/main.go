package main

import "fmt"

func main() {

	fmt.Println("Hi this is map examples")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of languages:", languages)
	fmt.Println("Js is the shortform of:", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// loops for maps

	for key, value := range languages {

		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
