package main

import "fmt"

func calculateSum(numbers []int) int {
sum := 0
for _, number := range numbers {
  sum += number
}

return sum

}




func main() {
    myNumbers := []int{1, 2, 3, 4, 5}
    total := calculateSum(myNumbers)
    fmt.Println(total) // Output: 15
}
