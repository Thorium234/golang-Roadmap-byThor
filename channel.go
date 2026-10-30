package main

import "fmt"

func main() {
    ch := make(chan int, 5)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    for value := range ch {
        fmt.Println("Value:", value)
    }
}
