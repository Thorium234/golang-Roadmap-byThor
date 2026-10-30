package main 

import "fmt"

func main(){
numbers := []int{10,30,40,50}
for index,value := range numbers{
	fmt.Println("Index:",index, "Value:",value)
}
}
//Iterating Over a String

//When iterating over a string, the for-range loop provides the index and the Unicode code point (rune) of each character.

/*

func main() {
    text := "Hello"
    for index, char := range text {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }
}





*/


