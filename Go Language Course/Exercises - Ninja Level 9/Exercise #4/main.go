package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0

	const gs = 20
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("count:", counter)
}
