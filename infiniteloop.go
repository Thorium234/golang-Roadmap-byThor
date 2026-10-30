package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("This loop runs forever...")
		time.Sleep(1 * time.Second) // Optional: keeps it from overwhelming your CPU
	}
}
