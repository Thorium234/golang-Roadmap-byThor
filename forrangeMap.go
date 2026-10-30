package main 
import "fmt"

func main(){
ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
for key, value := range ages{fmt.Println("Key:", key, "Value:", value)}
}
