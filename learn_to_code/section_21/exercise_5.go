package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementTheBoi() {
	atomic.AddInt32(&incrementBoi, 1)
	wg.Done()
}

var incrementBoi int32
var wg sync.WaitGroup

func main() {
	incrementBoi = 0
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go incrementTheBoi()
	}

	wg.Wait()
	fmt.Println(incrementBoi)
}
