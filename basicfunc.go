/*package main 

import "fmt"

func sayHello(){
	fmt.Println("hello thorium")
}

func main(){
sayHello()//output hello thorium
}
*/

package main
import "fmt"
func divide(c,d int)(quotient int,remainder int){
quotient = c/d
remainder = c%d
return
}
func add(a int, b int)int{
	return a+b
}
func main(){
sum := add(3,4)
fmt.Println("Sum:",sum) //output sum:7

q,r :=divide(10,3)
fmt.Println("Quotient: ",q,"Remainder: ",r)// Output: Quotient: 3 Remainder: 1
}
