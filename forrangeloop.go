package main

import "fmt"

func main(){
arr := []int{1,2,3,4,5}
for index,value := range arr{
fmt.Println("Index:",index,"Value:",value)
}
}
