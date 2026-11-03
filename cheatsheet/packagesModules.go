// Create a new module (in project folder)
// go mod init mymodule

// Add a dependency
// go get github.com/google/uuid

// go.mod file will track module path and dependencies

// Package example
package mathutil

// Exported function (capitalized)
func Add(a, b int) int {
    return a + b
}

// Using a package
import (
    "fmt"
    "mymodule/mathutil"
)

sum := mathutil.Add(3, 5)
fmt.Println(sum) // 8

// Alias import
import m "mymodule/mathutil"
result := m.Add(2, 4)

// Blank import (import only for side effects)
// import _ "example.com/some/package"

// List all modules and versions
// go list -m all

// Tidy up dependencies
// go mod tidy
