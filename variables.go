var a int //basic declaration
var b int = 10 //declaration with initialization
c := 22 //type interface
var d,e,f int //multiple variable declaration

package main  

import "fmt"

func main(){
var x int = 100 //x is accessible within the scope of main function
fmt.Println(x)

if true{
y := 200 // y is only accessible within this if block
fmt.Println(y)
}
//fmt.Println(y) //this will couse  error

//contants 
const pi = 3.14
const greet = "hello"
fmt.Println(pi)
fmt.Println(greet)
}
