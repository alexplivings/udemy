package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 5
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	go func() {
		c <- 6
		close(c)
	}()

	v, ok = <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}
