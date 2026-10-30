package main

import "fmt"

func main() {
    var x interface{} = 10

    switch v := x.(type) {
    case int:
        fmt.Println("x is an int") // Output: x is an int
    case string:
        fmt.Println("x is a string")
    default:
        fmt.Println("Unknown type")
    }
}
