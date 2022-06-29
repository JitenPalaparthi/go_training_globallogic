package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mx)
	}
	wg.Wait()
	fmt.Println("The final value of X", x)
}

var x = 0

func increment(wg *sync.WaitGroup, mx *sync.Mutex) {
	mx.Lock()
	x = x + 1
	mx.Unlock()
	wg.Done()
}

// The same task do it with channel
