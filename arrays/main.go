package main

import "fmt"

func main() {

	fmt.Println("This is array:")

	var fruitsList [4]string

	fruitsList[0] = "apple"
	fruitsList[1] = "mango"
	fruitsList[2] = "guava"

	fmt.Println("fruit list is:", fruitsList)
	fmt.Println("fruit list is:", len(fruitsList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list is:", vegList)

}
