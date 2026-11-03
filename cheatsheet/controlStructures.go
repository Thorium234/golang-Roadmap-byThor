// --- If / Else ---
x := 10
if x > 5 {
    fmt.Println("x > 5")
} else if x == 5 {
    fmt.Println("x == 5")
} else {
    fmt.Println("x < 5")
}

// --- Switch ---
day := 3
switch day {
case 1:
    fmt.Println("Monday")
case 2, 3:
    fmt.Println("Tue or Wed")
default:
    fmt.Println("Other day")
}

// Expression-less switch
num := 15
switch {
case num%2 == 0:
    fmt.Println("Even")
default:
    fmt.Println("Odd")
}

// --- For Loops ---
// Classic
for i := 0; i < 3; i++ {
    fmt.Println(i)
}

// While-style
j := 0
for j < 3 {
    fmt.Println(j)
    j++
}

// Infinite loop
// for { fmt.Println("loop") }

// --- Range ---
// Slice
nums := []int{10, 20, 30}
for idx, val := range nums {
    fmt.Println(idx, val)
}

// Map
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Println(k, v)
}

// String (runes)
for i, r := range "Go" {
    fmt.Printf("%d:%c\n", i, r)
}

// --- Break / Continue ---
for i := 0; i < 5; i++ {
    if i == 2 { continue }
    if i == 4 { break }
    fmt.Println(i)
}
