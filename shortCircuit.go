package main

import "fmt"

func main() {
    var a bool = false
    var b bool = true

    // The second condition (b) is not evaluated because the first condition (a) is false
    if a && b {
        fmt.Println("Both a and b are true")
    } else {
        fmt.Println("Either a or b is false") // Output: Either a or b is false
    }

    // The second condition (b) is not evaluated because the first condition (a) is true
    if a || b {
        fmt.Println("At least one of a or b is true") // Output: At least one of a or b is true
    } else {
        fmt.Println("Both a and b are false")
    }
}
