package main 
import "fmt"

func main(){
	var a int = 10
	var b int = 20

	if a > 5 {
	   if b > 15 {
		fmt.Println("a is greatort than 5 and b is greater that 15")
		}else{
			fmt.Println("a is greater that 5 but b is not greater than 15")			
			}
else{
fmt.Println("a is greater than 5 but b is not greater than 15 ")
}else{
fmt.Println("a is not greater than 5")
}		
}
}
