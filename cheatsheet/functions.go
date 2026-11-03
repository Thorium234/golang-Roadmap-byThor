// Basic function
func add(a int, b int) int {
    return a + b
}
sum := add(3, 5) // 8

// Multiple return values
func divide(a, b int) (int, int) {
    q := a / b
    r := a % b
    return q, r
}
q, r := divide(10, 3) // q=3, r=1

// Named return values
func rectStats(w, h int) (area int, perimeter int) {
    area = w * h
    perimeter = 2 * (w + h)
    return
}
a, p := rectStats(4, 5) // area=20, perimeter=18

// Variadic function
func sumAll(nums ...int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}
total := sumAll(1, 2, 3, 4) // 10

// Function as value
f := add
result := f(10, 20) // 30

// Anonymous function
greet := func(name string) string {
    return "Hello " + name
}
msg := greet("Go") // Hello Go
