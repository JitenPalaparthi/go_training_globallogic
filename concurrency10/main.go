package main

import (
	"fmt"
)

func main() {

	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)
	// select is used only with channels .
	// If more than one channel is emitting data, to accept the data from one of them
	// can use select
outer:
	for {
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			break outer
		case s2 := <-output2:
			fmt.Println(s2)
			break outer
		default:
			fmt.Println("Nothing happned")
		}
	}
}

func server1(ch chan string) {
	// ms := int64(rand.Intn(10))
	// time.Sleep(time.Millisecond * time.Duration(ms))
	ch <- "From Server1"
}

func server2(ch chan string) {
	// ms := int64(rand.Intn(11))
	// time.Sleep(time.Millisecond * time.Duration(ms))
	ch <- "From Server2"
}
