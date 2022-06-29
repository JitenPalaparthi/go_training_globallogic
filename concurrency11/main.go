package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	//wg.Add(10)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go greet(&wg)
	}

	fmt.Println("Hello World from main")
	wg.Wait()
}

func greet(wg *sync.WaitGroup) {
	fmt.Println("Hello world from greet function")
	wg.Done()
}
