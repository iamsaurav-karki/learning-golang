package main

import "fmt"

func main() {
	fmt.Println("Variables:")
	var username string = "saurav Karki"
	fmt.Println("variable:username is: ", username)
	fmt.Printf("variable of type : %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)

	fmt.Println("variable:username is: ", isLoggedin)
	fmt.Printf("variable of type : %T \n", isLoggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable of type: %T \n", smallval)

	var largeval uint64 = 2343324
	fmt.Println(largeval)
	fmt.Printf("variable of type: %T \n", largeval)

	var smallfloat float32 = 2343324.23432424234
	fmt.Println(smallfloat)
	fmt.Printf("variable of type: %T \n", smallfloat)

	// default value and some alias

	var novalue int
	fmt.Println(novalue)

	// no var style
	name := "saurav karki"
	fmt.Println(name)

}
