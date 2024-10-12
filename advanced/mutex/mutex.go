package main

import (
	"fmt"
	"sync"
)

// mutex is used to synchronize access to shared resources.
// it is used to prevent race conditions.
// a race condition occurs when two or more goroutines access the same variable concurrently, and at least one of the accesses is a write.


type Post struct {
	mu    sync.Mutex
	likes int
}


func (p *Post) AddLike(wg *  sync.WaitGroup) {
	defer func () {
		wg.Done()
	}()
	p.mu.Lock()
	p.likes++
	p.mu.Unlock()
}

func main() {
	post := Post{}
    var wg sync.WaitGroup
	for i := 0; i<100; i++ {
		 wg.Add(1)
		go post.AddLike(&wg)
	}
	wg.Wait()
	fmt.Println(post.likes)
}
