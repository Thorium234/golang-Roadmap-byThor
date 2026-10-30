package main 

import "fmt"

func main(){
mymap := map[string]int{"one":1,"two":2,"three":3}
for key, value := range mymap{
fmt.Println("Key: ",key, "Value: ",value)
}


}
