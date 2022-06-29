package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var val chan int
	fmt.Println("Receive rand int")
	val = receiveRand() // This retunrs the channel

	go func() {
		for {
			v, ok := <-val
			if ok {
				fmt.Println(v)
				close(val)
				break
			}
		}
	}()

	time.Sleep(time.Millisecond * 1)
}

func receiveRand() chan int {
	r := make(chan int, 1)
	go func() {
		r <- rand.Intn(1000)
	}()
	return r
}
