package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			delay := rand.Intn(100)
			time.Sleep(time.Millisecond * time.Duration(delay))
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
