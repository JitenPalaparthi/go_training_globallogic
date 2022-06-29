package main

import (
	"fmt"
	"time"
)

func main() {
	go greet()

	fn := func() {
		for i := 1; i <= 10; i++ {
			go fmt.Println("Hello Anonymous function")
		}
	}

	go fn()
	// anonymous function
	go func(si, ei int) {
		for i := si; i <= ei; i++ {
			go greetI(i)
		}
	}(1, 100)

	time.Sleep(time.Millisecond * 10)
	fmt.Println("Hello World")
}

func greet() {
	go fmt.Println("Hello Go routine")

	go func() {
		fmt.Println("----> Hello Go routine <---------")
	}()
}

func greetI(i int) {
	fmt.Println("Hello Go routine--->", i)
}

// go routine can be thought of a simple thread. But it not a thread what we learnt in java or .net
// go routines are scheduled and maintianed by go runtime not by the OS
// 1. go is a keyword to create a go routine.Any function can be a go routine
// 2. Main is a also a go routinego
// 3. By default , no go routine waits for other go routine to complete its task. It is the responsibility of the programer to check or perform opetations
// 4. There is no guarantee abt the order of execution
