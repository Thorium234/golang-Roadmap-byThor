// Variable declaration
x := 42
p := &x        // pointer to x
*p = 100       // change x through pointer

// Function parameter (pass by reference)
func setTo200(n *int) {
    *n = 200
}
setTo200(&x)   // x becomes 200

// Function return (return pointer)
func createPointer() *int {
    v := 300
    return &v
}
ptr := createPointer()
fmt.Println(*ptr) // 300
