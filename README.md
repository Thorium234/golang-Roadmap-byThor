

---

### **Slide 1: Program Structure & Basic Types**

#### **Minimal Executable Program**
* Executable Go applications must start with a package declaration named `main` and execute starting from the `main()` function [1, 10, 11].
* Code from external packages is included using the `import` keyword [12].

```go
package main 

import "fmt" 

// Main entry point
func main() {     
    fmt.Println("Hello World") 
} 
``` [1]

#### **Built-in Primitive Data Types**
* **Numbers:** Integer types include signed (`int8` to `int64`) and unsigned (`uint8` to `uint64`), while floating-point types include `float32` and `float64` [13, 14].
* **Strings:** Sequences of bytes supporting length checking via `len()`, zero-based index access, and string concatenation using `+` [15-17].
* **Booleans:** Logical values (`true`/`false`) evaluated with `&&` (AND), `||` (OR), and `!` (NOT) [18].

```go
// Numbers
fmt.Println("1 + 1 =", 1 + 1)         // int [2]
fmt.Println("1 + 1 =", 1.0 + 1.0)     // float64 [15]

// Strings
fmt.Println(len("Hello World"))        // returns 11 [17]
fmt.Println("Hello World"[19])          // returns byte representation '101' [17]
fmt.Println("Hello " + "World")        // concatenation [17]

// Booleans
fmt.Println(true && false)             // returns false [18]
``` [2, 15, 17, 18]

---

### **Slide 2: Variables & Control Flow**

#### **Variable Declarations & Scope**
* Variables are declared explicitly using `var` or inferred short declaration (`:=`) [3, 20].
* Constants are declared with `const` and cannot be modified after assignment [21, 22].
* Variables are block-scoped within the nearest enclosing curly braces `{}` [21].

```go
// Long declaration
var x string = "Hello World" [3]

// Short variable declaration (type inferred)
x := "Hello World" [20]

// Grouped variable declaration
var (
    a = 5
    b = 10
    c = 15
) [23]

// Constant definition
const y string = "Hello World" [22]

// Reading user input
var input float64
fmt.Scanf("%f", &input) [23]
``` [3, 20, 22, 23]

#### **Control Structures**
* Go uses `for` as its sole looping construct, replacing `while` and `foreach` [4, 24].
* Conditionals use `if`/`else` and `switch` statements [25, 26].

```go
// Standard for loop
for i := 1; i <= 10; i++ {
    fmt.Println(i)
} [24]

// If - Else branching
if i % 2 == 0 {
    fmt.Println(i, "even")
} else {
    fmt.Println(i, "odd")
} [27]

// Switch branching
switch i {
case 0: 
    fmt.Println("Zero")
case 1: 
    fmt.Println("One")
default: 
    fmt.Println("Unknown Number")
} [26]
``` [24, 26, 27]

---

### **Slide 3: Data Structures (Arrays, Slices & Maps)**

#### **Arrays vs. Slices**
* **Arrays:** Fixed-length sequences of a single type [28].
* **Slices:** Resizable segments of underlying arrays initialized with `make([]T, len, cap)` or slice operations `[low:high]` [5, 29].

```go
// Fixed array
x := [30]float64{ 98, 93, 77, 82, 83 } [31]

// Range loop (blank identifier _ ignores the index)
total := 0.0
for _, value := range x {
    total += value
} [31]

// Slice creation using make(type, length, capacity)
s := make([]float64, 5, 10) [5, 29]

// Slicing an existing array/slice
sub := x[0:5] [29]

// Append and Copy helper functions
slice1 := []int{1, 2, 3} [32]
slice2 := append(slice1, 4, 5)     // slice2 = [19, 30, 33-35] [32]
target := make([]int, 2)
copy(target, slice1)               // target = [19, 33] [36]
``` [5, 29, 31, 32, 36]

#### **Maps (Associative Arrays)**
* Unordered key-value hash tables created via `make(map[KeyType]ValueType)` [36, 37].
* Lookups return two values: the value and a boolean (`ok`) indicating key existence [38, 39].

```go
// Map initialization
elements := make(map[string]string) [40]
elements["H"] = "Hydrogen" [40]

// Map item deletion
delete(elements, "H") [41]

// Safe key existence lookup
if name, ok := elements["Un"]; ok {
    fmt.Println(name, ok)
} [39]
``` [39-41]

---

### **Slide 4: Functions, Closures & Resource Management**

#### **Functions & Return Values**
* Functions start with `func`, support multiple return values, and accept variadic parameters (`...T`) [6, 42, 43].

```go
// Multiple return values
func f() (int, int) {
    return 5, 6
} [6]

// Variadic function
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
} [43]
``` [6, 43]

#### **Closures & Recursion**
* Functions are first-class types capable of forming closures by referencing outer local variables [44, 45].

```go
// State-preserving function generator (Closure)
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
} [46]

// Recursive function
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
} [47]
``` [46, 47]

#### **Defer, Panic & Recover**
* `defer` delays execution of a function call until the outer function completes [48].
* `panic` triggers a runtime error, which can be caught using `recover` inside a deferred call [49, 50].

```go
// Resource cleanup pattern
f, _ := os.Open("filename")
defer f.Close() [51]

// Catching a panic runtime error
defer func() {
    str := recover()
    fmt.Println(str)
}() [50]
panic("PANIC") [50]
``` [50, 51]

---

### **Slide 5: Pointers, Structs & Interfaces**

#### **Pointers**
* Pointers store memory addresses of values using `*` (type/dereference) and `&` (address-of) [7, 52, 53].

```go
func zero(xPtr *int) {
    *xPtr = 0 // dereference and modify value [7, 52]
}

func main() {
    x := 5
    zero(&x)  // pass address of x [7, 53]
    
    ptr := new(int) // allocate memory returning *int [53]
}
``` [7, 52, 53]

#### **Structs & Receivers (Methods)**
* Custom structures contain named fields; methods attach to structs using receivers [8, 54].
* Composition ("is-a" relationship) is achieved using **embedded anonymous fields** [55].

```go
type Circle struct {
    x, y, r float64 [56]
}

// Method with pointer receiver
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r [54]
}

// Embedded Struct Composition
type Person struct {
    Name string
} [57]

type Android struct {
    Person // Anonymous embedded field [55]
    Model  string
} [55]
``` [54-57]

#### **Interfaces**
* Interfaces define method sets satisfied implicitly by types implementing those methods [58].

```go
type Shape interface {
    area() float64
} [58]

// Interface usage in functions
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
} [59]
``` [58, 59]

---

### **Slide 6: Concurrency (Goroutines & Channels)**

#### **Goroutines**
* Lightweight concurrent execution threads started with the `go` keyword [9, 60].

```go
func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
    }
} [9]

func main() {
    go f(0) // Executes concurrently [9]
}
``` [9]

#### **Channels & Multiplexing**
* Channels synchronize execution and transmit values between goroutines using `<-` [61, 62].

```go
// Channel creation
c := make(chan string) [61]

// Send & Receive
c <- "ping"    // Send [62]
msg := <- c    // Receive [62]

// Channel Direction Constraints
func pinger(c chan<- string) // Send-only [63]
func printer(c <-chan string) // Receive-only [63]

// Select Multiplexing with Timeout
select {
case msg1 := <- c1:
    fmt.Println("Message 1", msg1) [64]
case <- time.After(time.Second):
    fmt.Println("timeout") [64]
default:
    fmt.Println("nothing ready") [65]
} [64, 65]

// Buffered Channel (Capacity of 1)
bufChan := make(chan int, 1) [65]
``` [61-65]

---

### **Slide 7: Packages, Testing & Core Libraries**

#### **Packages & Visibility**
* Capitalized identifiers are **exported** (public outside package), while lower-case identifiers remain package-private [66, 67].

```go
package math

// Exported function (Capitalized)
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
} [68]
``` [66, 68]

#### **Unit Testing**
* Tests live in files ending with `_test.go` and run via `go test` [69, 70].

```go
package math

import "testing"

func TestAverage(t *testing.T) {
    v := Average([]float64{1, 2})
    if v != 1.5 {
        t.Error("Expected 1.5, got ", v)
    }
} [70]
``` [70]

#### **Essential Standard Library Snippets**

```go
// 1. File Reading (io/ioutil)
bs, err := ioutil.ReadFile("test.txt") [71]

// 2. HTTP Web Server (net/http)
http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "Hello World!")
}) [72]
http.ListenAndServe(":9000", nil) [72]

// 3. Thread Synchronization Mutex (sync)
m := new(sync.Mutex) [73]
m.Lock() [73]
// Protected critical section [73]
m.Unlock() [73]
``` [71-73]
