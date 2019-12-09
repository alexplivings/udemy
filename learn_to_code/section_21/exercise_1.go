package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func oneFunc() {
	fmt.Println("Doing something in oneFunc")
	time.Sleep(time.Second * 10)
	fmt.Println("Done with oneFunc")
	wg.Done()
}

func twoFunc() {
	fmt.Println("Doing something in twoFunc")
	time.Sleep(time.Second * 5)
	fmt.Println("In the middle of something in twoFunc")
	time.Sleep(time.Second * 8)
	fmt.Println("Done with twoFunc")
	wg.Done()
}

func main() {

	wg.Add(2)

	go oneFunc()

	go twoFunc()

	wg.Wait()
}
