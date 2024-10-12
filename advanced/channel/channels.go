package main

import (
	"fmt"
)

// channels are used to communicate between goroutines.
// when you use channels or waitgroups to synchronize goroutines, when you work with multiple goroutines, you need to use waitgroups and when only work with one goroutines then use channels.

// if you don't close the channel, the program will deadlock.
// you can use the built-in close function to close a channel.

// you can use the range keyword to iterate over a channel until it is closed.

// when you send a value to a channel, the channel will block until the value is received.
// when you receive a value from a channel, the channel will block until the value is sent.

func process(msg chan string) {
	fmt.Println("Processing", <-msg)
}

func main() {

	messages := make(chan string)

	go process(messages)
	messages <- "hello world"
	// close(messages)
	fmt.Println("channels")

	// time.Sleep(time.Second * 1)
}