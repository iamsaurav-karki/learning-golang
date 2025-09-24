package main

import "fmt"

func main() {
	fmt.Println("Arrays:")
	var number [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("number=", number)

	var fruits [5]string = [5]string{"mango", "apple", "banana", "grapes", "pineapple"}
	fmt.Println("Fruits=", fruits)
}
