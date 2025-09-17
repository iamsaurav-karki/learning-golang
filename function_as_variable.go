package main
import "fmt"

func main(){
    var v func(int) int          // Step 1: Declare a variable 'v' of type 'function'
    v = func(x int) int {        // Step 2: Assign an anonymous function to 'v'
        return x * x
    }
    result := v(5)               // Step 3: Call 'v' with argument 5
    fmt.Println("result=", result) // Step 4: Print result
}

