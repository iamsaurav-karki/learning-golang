package main

import "fmt"

func main() {

	fmt.Println("This is loops")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// one way

	for d := 0; d < len(days); d++ {
		fmt.Print(days[d])
	}

	// second way

	for i := range days {
		fmt.Println(days[i])

	}

	// Third way

	for index, day := range days {

		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// Fourth way

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			break
		}

		if rougueValue == 3 {
			rougueValue++
			continue
		}

		if rougueValue == 2 {
			goto lco
		}

		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at sauravverse.xyz")

}
