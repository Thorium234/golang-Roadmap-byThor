// Struct definition
type Person struct {
    Name string
    Age  int
}

// Create struct instance
p1 := Person{"Alice", 30}
p2 := Person{Name: "Bob"} // Age = 0 by default

// Access fields
fmt.Println(p1.Name) // Alice
fmt.Println(p1.Age)  // 30

// Pointer to struct
p3 := &Person{"Charlie", 25}
fmt.Println(p3.Name) // Charlie

// Method with value receiver
func (p Person) Greet() string {
    return "Hello, " + p.Name
}
msg := p1.Greet() // Hello, Alice

// Method with pointer receiver (can modify struct)
func (p *Person) HaveBirthday() {
    p.Age++
}
p1.HaveBirthday()
fmt.Println(p1.Age) // 31

// Embedding (like inheritance)
type Employee struct {
    Person
    Position string
}
e := Employee{Person{"Diana", 28}, "Engineer"}
fmt.Println(e.Name)     // Diana (from embedded Person)
fmt.Println(e.Position) // Engineer
