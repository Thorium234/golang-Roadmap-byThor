package main 
import "fmt"

func sum(numbers ...int) int{
total := 0
for _, number := range numbers{
total += number
}
return total

}
func sum2D(matrix [3][5]int) int {
grandTotal := 0
for _,row := range matrix{
grandTotal += sum(row[:]...)
}
return grandTotal
}
//var a [5] int

func main(){
//1D arr
a := [5]int {1,2,3,4}

//2D arr
b := [3][5]int{
{1,2,3,4,5},{2,3,4,5,6},{6,7,8,9,1},
}
//result := sum(a[0],[1])
//fmt.Println("arry sum:",sum(a))//output 10
fmt.Println("1D-arry sum:",sum(a[:]...))
fmt.Println("2D-array sum:",sum2D(b))
fmt.Println("sum:",sum(1,2,3,4))//output sum:10
//fmt.Println("Sum of two elements:", result) // Output: 3

// a[1:3] grabs elements from index 1 up to (but excluding) index 3 -> {2, 3}
result := sum(a[1:3]...)
fmt.Println("Sum of sub-slice:", result) // Output: 5
// Let's sum the number 1 (row 0, col 0) and the number 9 (row 2, col 3) from your matrix 'b'
result := sum(b[0][0], b[2][3])
fmt.Println("Sum of two matrix elements:", result) // Output: 10
// Take row 1 ({2, 3, 4, 5}) and slice from index 0 up to (but excluding) index 2 -> {2, 3}
result := sum(b[1][0:2]...)
fmt.Println("Sum of row sub-slice:", result) // Output: 5


}
