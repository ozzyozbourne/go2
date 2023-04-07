package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 3; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("The value of i -> %d\n", i)
		}()
		wg.Wait()
	}
}
