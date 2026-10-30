package main

import "fmt"

func main() {
    var a int = 5
    var b int = 3
    var c int = 8

    var result bool = (a > b) && (c > a)
    fmt.Println("(a > b) && (c > a):", result) // Output: (a > b) && (c > a): true
}
