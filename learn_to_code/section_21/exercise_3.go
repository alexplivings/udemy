package main

import (
	"fmt"
	"runtime"
	"sync"
)

func incrementTheBoi() {
	x := incrementBoi
	runtime.Gosched()
	x++
	fmt.Println("Setting incrementBoi to", x)
	incrementBoi = x
	wg.Done()
}

var incrementBoi int
var wg sync.WaitGroup

func main() {
	incrementBoi = 0
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go incrementTheBoi()
	}

	wg.Wait()
}
