package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/* Solution 1
	var mu sync.Mutex
	count := 0
	*/

	// count := int64(0)
	var count int64

	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				/* Solution 1
				mu.Lock()
				count++
				mu.Unlock()
				*/
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// go run -race .
