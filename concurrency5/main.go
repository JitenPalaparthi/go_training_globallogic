package main

import "fmt"

var done chan bool

func main() {
	done = make(chan bool)
	go greet()
	<-done // This is blocked untile a value is sent to done so that it can be received here
	fmt.Println("Hello Folks! This is from main")
}

func greet() {
	fmt.Println("Hello World!  This is from greet go routine")
	done <- true
}
