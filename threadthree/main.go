package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type person struct {
	string
}

func (p *person) display() {
	fmt.Printf("%#v\n", p)
}

func main() {
	var p person
	p.display()
	pn := new(person)
	pn.display()
	runtime.GOMAXPROCS(0)
	var counter int
	wg := new(sync.WaitGroup)
	wg.Add(100)
	mutex := new(sync.Mutex)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			runtime.Gosched()
			mutex.Lock()
			v := counter
			v++
			counter = v
			fmt.Printf("No of CSPs -> %-15d Counter Val -> %-15d Actual V Val -> %d\n",
				runtime.NumGoroutine(), counter, v)
			mutex.Unlock()

		}()
	}
	wg.Wait()
	fmt.Printf("Counter Val -> %d\n", counter)
}
