package main
import "fmt"
func main(){
	fmt.Println("slices:")
	person := []string{"Saurav","samyog","sulav","hari"}  // slice of strings
	fmt.Println("p=", person)
	displayfruits()
}

// slice is more flexible than array so it is used more often then array..

func displayfruits(){
	fmt.Println("fruits:")
	fruits := []string{"mango","banana","apple","guava","watermelon","pear","peach"}
	fmt.Println("fruits=",fruits)
}

// In case of array we need to define specific array size and not allowed to add more than the defined size but in case of slices we need not to define the size explicitly and can add any items we want.


