package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("NO_OF_CSPS: %d\n", runtime.NumGoroutine())
	runtime.GOMAXPROCS(1)
	wg.Add(4)
	go foo()
	go bar()
	go func() {
		foo()
		bar()
	}()
	fmt.Printf("GOARCH: %s\nGOOS: %s\nMAXPROCS: %d\nNO_OF_CSPS: %d\nN0_OF_CORE_FOR_SCHEDULING: %d\n",
		runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.NumGoroutine(), runtime.GOMAXPROCS(0))
	wg.Wait()
	fmt.Printf("NO_OF_CSPS: %d\n", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("foo: %d\n", 1)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Printf("bar: %d\n", 1)
	}
	wg.Done()
}
