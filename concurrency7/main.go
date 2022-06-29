package main

import (
	"fmt"
)

func main() {
	data := make(chan string, 10)
	data <- "Hello"
	data <- "World!"
	fmt.Println(<-data)
	data <- "How are you doing?"
	fmt.Println(<-data)
	fmt.Println(<-data)
	// Function to return a channel
	//var val chan int
}

// If the size of the channel is zero then it is called as unbuffered channel
// If the size of the channel is not zero then it is called as buffered channel
// The sender and receiver will not be blocked until the buffer is full
