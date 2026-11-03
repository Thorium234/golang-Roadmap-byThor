import (
    "errors"
    "fmt"
)

// Built-in error type
err := errors.New("something went wrong")
if err != nil {
    fmt.Println(err)
}

// Functions returning error
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}

// fmt.Errorf for formatted errors
name := ""
if name == "" {
    err := fmt.Errorf("invalid name: %v", name)
    fmt.Println(err)
}

// Panic: stops program immediately
// Used for unrecoverable errors
// panic("fatal error")

// Recover: used inside defer to catch panic
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    if b == 0 {
        panic("divide by zero")
    }
    return a / b
}

fmt.Println(safeDivide(10, 2)) // 5
fmt.Println(safeDivide(10, 0)) // recovered
