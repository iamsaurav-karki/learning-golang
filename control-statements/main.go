package main

import "fmt"

func main() {

	fmt.Println("This is the example of control statements.")

	loginCount := 10
	var result string

	if loginCount < 10 {

		result = "Regular User"
	} else if loginCount > 10 {

		result = "watch out"
	} else {

		result = "Exactly 10 login counts"
	}
	fmt.Println(result)

	if loginCount%2 == 0 {

		fmt.Println("The number is even.")
	} else {

		fmt.Println("The number is odd.")
	}

	if num := 3; num < 10 {

		fmt.Println("Num is less than 10")

	} else {

		fmt.Println("The number is not less than 10")
	}

	// if err != nil {

	// }

}
