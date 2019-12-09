package main

import (
	"fmt"
	"runtime"
	"sync"
)

func incrementTheBoi() {
	mu.Lock()
	x := incrementBoi
	runtime.Gosched()
	x++
	fmt.Println("Setting incrementBoi to", x)
	incrementBoi = x
	mu.Unlock()
	wg.Done()
}

var incrementBoi int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	incrementBoi = 0
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go incrementTheBoi()
	}

	wg.Wait()

	fmt.Println(incrementBoi)
}
