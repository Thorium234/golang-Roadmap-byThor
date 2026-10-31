package main 
import "fmt"

func sum(numbers ...int) int{
total := 0
for _, number := range numbers{
total += number
}
return total

}

//var a [5] int

func main(){
a := [5]int {1,2,3,4}
//fmt.Println("arry sum:",sum(a))//output 10
fmt.Println("arry sum:",sum(a[:]...))
fmt.Println("sum:",sum(1,2,3,4))//output sum:10
}
