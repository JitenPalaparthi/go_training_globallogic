package main

import (
	"fmt"
	"time"
)

// Do not communicate by sharing memory; Instead , share memory by communicating
var count int = 0

func main() {
	var ch chan int
	ch = make(chan int)
	// 1- chan is a keyword to create a channel.
	// 2- Followed by chan the datatype to be specified
	// 3- Datatype can be any valid datatype

	go func() {
		for i := 1; i <= 1000; i++ {
			go Increment()
			go Decrement()
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("The final Count-->", count)
}

func Increment() {
	count++
	fmt.Println("Increment-->", count)
}

func Decrement() {
	count--
	fmt.Println("Decrement-->", count)
}
