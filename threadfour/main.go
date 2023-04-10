package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(0)
	counter := new(int64)
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			runtime.Gosched()
			atomic.AddInt64(counter, 1)
			fmt.Printf("No of CSPs -> %-15d Counter Val -> %-15d\n",
				runtime.NumGoroutine(), atomic.LoadInt64(counter))
		}()
	}
	wg.Wait()
	fmt.Printf("Counter Val -> %d\n", *counter)
}
