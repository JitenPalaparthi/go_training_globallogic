package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // chan is a ref becase it is created with make function
	go greet(done)

	data := make(chan string)
	go publish(data)
	go subscribe(data)

	<-done // This is blocked untile a value is sent to done so that it can be received here
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Hello Folks! This is from main")
}

// there is send only channel and receive only channel

func greet(ok chan<- bool) { // send only channel
	//b := <-ok receiving a value from channel is not possible here becasue it is a send only channel
	fmt.Println("Hello World!  This is from greet go routine")
	ok <- true
}

func publish(send chan<- string) { // send only channel
	//b := <-ok receiving a value from channel is not possible here becasue it is a send only channel
	fmt.Println("I am a publisher.. publishing data!")
	send <- "Hello Subscriber, from Publisher"
}

func subscribe(receive <-chan string) { // receive only channel
	//b := <-ok receiving a value from channel is not possible here becasue it is a send only channel
	data := <-receive
	//receive <- "Hello" This does not work because it is receive only channel
	fmt.Println(data, ".I am a subscriber received data")
}
