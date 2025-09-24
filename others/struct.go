package main 
import "fmt"

func main(){
	fmt.Println("Struct Example:")
	displaydetails()
}

func displaydetails(){
	type Person struct {
		Name string
		Age int
		Address string
	}
	var q Person = Person{"Haris", 30, "Dharan"}
	fmt.Println("q=", q)
	fmt.Println("Name=", q.Name)
	fmt.Println("Age=", q.Age)
	fmt.Println("Address=", q.Address)
}
