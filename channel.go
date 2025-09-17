package main

import "fmt"

func main(){

	fmt.Println("Channel example:")
	u := make(chan int) // unbuffered channel of integers
	fmt.Println("u=", u)
}
