package main

import "fmt"

func main(){
	fmt.Println("Pointers Example")
	var s int = 20
	var t *int = &s // pointer to s  & is used to get the memory address of s. 
	fmt.Println("value=", s)
	fmt.Println("Address=",t)
	fmt.Println("ValuebyAddress=",*t)


	var name string = "Hari"
	var nameaddress *string = &name 

	fmt.Println("value=", name)
	fmt.Println("Address=", nameaddress)
	fmt.Println("ValuebyAddress=",*nameaddress)
}
