package main

import (
	"fmt"
	"sync"
)

// goroutine is a lightweight thread managed by the Go runtime.


func task (i int) {
	fmt.Println("running task",i)
}

func main() {

	var wg sync.WaitGroup

	// go keyword  use to run the function as a goroutine.
	for i:= 0; i<10; i++ {

		wg.Add(1)
		// anonymous function
		go func(i int, wg * sync.WaitGroup) {
			fmt.Println("running task",i)
			wg.Done()
		}(i, &wg)
	}
	
	// wait for all goroutines to finish
	// time.Sleep(time.Second * 2) 

	// We can't give time.sleep to main function because it will block the main function and the program will not exit. we also don't know how much time to wait.

	
	// we use waitgroup to wait for all goroutines to finish.

	wg.Wait()
	
}
