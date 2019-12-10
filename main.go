package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
	arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}
	//  ch := make(chan struct{})
	var wg sync.WaitGroup // number of working goroutines
	var mu sync.Mutex     // guards icons
	for _, f := range arr {
		mu.Lock()
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			fmt.Println("array item :: ", f)
			// ch <- struct{}{}
		}(f)
		mu.Unlock()
	}
	for range arr {
		wg.Wait()
		// <-ch
	}
}
