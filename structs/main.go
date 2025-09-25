package main

import "fmt"

func main() {

	fmt.Println("Hello this is the example for structs")

	details := User{"saurav", "sk@gmail.com", true, 29}
	fmt.Println(details)

	fmt.Printf("Saurav Details are: %+v\n", details)
	fmt.Printf("Name is %v and Email is %v \n.", details.Name, details.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
