package main
import "fmt"

// Employee struct represents an employee record

type Employee struct {
	Name string
	Position string
	Salary float64

}

// UpdateSalary updates the salary of the employee 

func UpdateSalary(emp *Employee, newSalary float64) {
	emp.Salary = newSalary
}

func main(){

	// Creating instance of employee
	emp := Employee{
		Name: "hari rai",
		Position: "Ml Engineer",
		Salary: 50000,
	}

	//Printing initial state 
	fmt.Println("Before update:", emp)

	//Updating Salary
	fmt.Println(&emp)
	UpdateSalary(&emp, 600000)

	//Printing updated state 

	fmt.Println("After Update:", emp)
}
