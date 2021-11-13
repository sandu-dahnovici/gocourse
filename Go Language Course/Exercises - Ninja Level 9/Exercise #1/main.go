package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("First : ", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Second : ", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 20; i++ {
			fmt.Println("Third : ", i)
		}
		wg.Done()
	}()
	wg.Wait()
}
