// --- Array ---
// Fixed size, elements of the same type
var arr [3]int = [3]int{1, 2, 3}
fmt.Println(arr[0]) // 1

// --- Slice ---
// Dynamic size, built on arrays (commonly used)
var sl []string = []string{"Go", "Rust", "Python"}
sl = append(sl, "JavaScript")
fmt.Println(sl[0]) // "Go"

// --- Map ---
// Key-value pairs (unordered)
var m map[string]int = map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(m["Alice"]) // 25
