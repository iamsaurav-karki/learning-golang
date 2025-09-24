package main

import "fmt"

func main() {

	fmt.Println("welcome to pointer") // pointer is nothing but a direct reference to a memory address.

	mynumber := 23

	var ptr = &mynumber

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The new value after adding 2 is:", mynumber)
}
