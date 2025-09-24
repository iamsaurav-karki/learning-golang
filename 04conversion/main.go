package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the ratings,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {

		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating:", numRating+1)
	}
}
