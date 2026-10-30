package main 

import "fmt"

func main(){
	var a int = 5
	var b int = 3 
	var sum int = a+b
	
	fmt.Println("Sum:",sum) //output 8
    var difference int = a - b
    fmt.Println("Difference:", difference) // Output: Difference: 2
    var product int = a * b
    fmt.Println("Product:", product) // Output: Product: 15
    var quotient int = a / b
    fmt.Println("Quotient:", quotient) // Output: Quotient: 2
    var remainder int = a % b
    fmt.Println("Remainder:", remainder) // Output: Remainder: 2
}
