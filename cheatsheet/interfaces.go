// Interface definition
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Struct implementing interface
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Another struct implementing interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

// Using the interface
var s Shape

s = Rectangle{Width: 4, Height: 5}
fmt.Println(s.Area())      // 20
fmt.Println(s.Perimeter()) // 18

s = Circle{Radius: 3}
fmt.Println(s.Area())      // 28.26
fmt.Println(s.Perimeter()) // 18.84

// Empty interface (can hold any type)
var any interface{}
any = "Hello"
any = 42
any = 3.14

// Type assertion
val, ok := any.(int)
if ok {
    fmt.Println("Integer:", val)
}

// Type switch
switch v := any.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown type")
}
