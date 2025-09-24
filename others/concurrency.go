package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()               // run concurrently
	time.Sleep(1 * time.Second) // wait for goroutine
	fmt.Println("Main function ends")
}
