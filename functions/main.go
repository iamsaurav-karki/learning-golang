package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Hi this is function")
}

func greeter() {

	addition()
	fmt.Println("Namaste from golang")
}

func addition() {

	result := adder(3, 5)
	sum := 2 + 3
	fmt.Println(sum)
	fmt.Println(result)

	proResult, mymsg := proAdder(2, 5, 7, 6)
	fmt.Println(proResult)
	fmt.Println(mymsg)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// Not allowed to write function inside functions

func proAdder(values ...int) (int, string) { // variatic functions.. can take any number of arguments.

	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi pro result fnx"

}
