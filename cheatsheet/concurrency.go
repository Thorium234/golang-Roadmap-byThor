// Goroutine (lightweight thread)
go fmt.Println("Hello from goroutine")

// Unbuffered channel (synchronizes send/receive)
ch := make(chan int)

// Send to channel
go func() {
    ch <- 42
}()

// Receive from channel
val := <-ch
fmt.Println(val) // 42

// Buffered channel
buf := make(chan string, 2)
buf <- "A"
buf <- "B"
fmt.Println(<-buf) // A
fmt.Println(<-buf) // B

// Closing channel
close(ch)
// after close, receive gives zero value
// val, ok := <-ch (ok=false if closed)

// Select with channels
c1 := make(chan string)
c2 := make(chan string)

go func() { c1 <- "one" }()
go func() { c2 <- "two" }()

select {
case msg1 := <-c1:
    fmt.Println("Received:", msg1)
case msg2 := <-c2:
    fmt.Println("Received:", msg2)
default:
    fmt.Println("No message")
}
