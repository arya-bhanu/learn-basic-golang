package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var meter counter
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for range 10000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			fmt.Printf("Meter before Add(): %d\n", meter.Get())
			meter.Add()
			fmt.Printf("Meter.Get(): %d\n", meter.Get())
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(meter.Get())
}

type counter struct {
	val int
}

func (c *counter) Get() int {
	return c.val
}

func (c *counter) Add() int {
	c.val++
	return c.val
}
