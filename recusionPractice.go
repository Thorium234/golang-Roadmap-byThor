package main

import "fmt"

func sumArray(arr []int, n int) int {
    if n <= 0 {
        return 0 // Base case
    }
    return arr[n-1] + sumArray(arr, n-1) // Recursive case
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println("Sum of array:", sumArray(arr, len(arr))) // Output: Sum of array: 15
}
