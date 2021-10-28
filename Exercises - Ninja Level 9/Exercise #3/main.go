package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0

	const gs = 20
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
