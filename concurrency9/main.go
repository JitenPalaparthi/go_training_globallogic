package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	//go func() {
	for v := range ch {
		fmt.Println(v)
	}
	//}()
	// The above range and this loop acheive the same thing
	// for {
	// 	val, ok := <-ch
	// 	if ok {
	// 		fmt.Println(val)
	// 	} else {
	// 		break
	// 	}
	// }

	//time.Sleep(time.Second * 1)
}
