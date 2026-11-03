// --- Arithmetic Operators ---
a, b := 10, 3
fmt.Println(a + b) // 13 (addition)
fmt.Println(a - b) // 7  (subtraction)
fmt.Println(a * b) // 30 (multiplication)
fmt.Println(a / b) // 3  (integer division)
fmt.Println(a % b) // 1  (modulus)

// --- Relational (Comparison) Operators ---
fmt.Println(a == b) // false
fmt.Println(a != b) // true
fmt.Println(a > b)  // true
fmt.Println(a < b)  // false
fmt.Println(a >= b) // true
fmt.Println(a <= b) // false

// --- Logical Operators ---
x, y := true, false
fmt.Println(x && y) // false (AND)
fmt.Println(x || y) // true  (OR)
fmt.Println(!x)     // false (NOT)

// --- Bitwise Operators ---
m, n := 6, 3         // 6 = 110, 3 = 011 in binary
fmt.Println(m & n)   // 2 (AND: 110 & 011 = 010)
fmt.Println(m | n)   // 7 (OR:  110 | 011 = 111)
fmt.Println(m ^ n)   // 5 (XOR: 110 ^ 011 = 101)
fmt.Println(m << 1)  // 12 (Left shift: 110 << 1 = 1100)
fmt.Println(m >> 1)  // 3  (Right shift: 110 >> 1 = 011)

// --- Assignment Operators ---
c := 5
c += 3  // c = c + 3 → 8
c -= 2  // c = c - 2 → 6
c *= 4  // c = c * 4 → 24
c /= 6  // c = c / 6 → 4
c %= 3  // c = c % 3 → 1

// --- Other Operators ---
// Address-of & Pointer Dereference
val := 42
ptr := &val        // address-of
fmt.Println(*ptr)  // dereference → 42
