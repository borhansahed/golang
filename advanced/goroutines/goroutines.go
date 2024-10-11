package main

import (
	"fmt"
	"time"
)

// goroutine is a lightweight thread managed by the Go runtime.


func task (i int) {
	fmt.Println("running task",i)
}

func main() {

	// go keyword  use to run the function as a goroutine.
	for i:= 0; i<10; i++ {
		// anonymous function
		go func(i int) {
			fmt.Println("running task",i)
		}(i)
	}
	
	// wait for all goroutines to finish
	time.Sleep(time.Second * 2)
}
