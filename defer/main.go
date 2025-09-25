package main

import "fmt"

func main() {
	defer fmt.Println("World") // defer moves the execution to last. just before the closing curley bracket
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Hello")
	myDefer()
}

func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
