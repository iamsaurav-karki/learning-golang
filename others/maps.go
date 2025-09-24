package main
import "fmt"

func main(){
	fmt.Println("Maps Example:")
	persons := map[string]int{"Alice": 30, "Bob": 22} // maps with string keys and integer values
	fmt.Println("personsdetails=", persons)
	fmt.Println("Alice=", persons["Alice"])
	fmt.Println("Bob=", persons["Bob"])
}
