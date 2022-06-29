package main

import "fmt"

// Do not communicate by sharing memory; Instead , share memory by communicating
func main() {
	var ch chan int
	ch = make(chan int) // unbuffered channel. The capacity is 0
	// 1- chan is a keyword to create a channel.
	// 2- Followed by chan the datatype to be specified
	// 3- Datatype can be any valid datatype
	// 4- Data to be sent or received to and from channel
	// 5- Send data to the channel ch <- 100
	// 6- Receive data from the channel data:= <-ch
	// 7- Channels are two types. They are buffered and unbuffered
	// 8- Once the data is sent to the channel, until that data is received
	// 	  from the channel the sender and the receiver go routines are blocked.
	go func() {
		fmt.Println("This is blocked until the values is received")
		r := <-ch
		fmt.Println("Received data from channel is", r)
	}()
	//func() {
	ch <- 1
	//}()

}
