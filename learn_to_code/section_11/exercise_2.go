package main

import "fmt"

func main() {
	values := []int{42, 420, 1337, 1, 5, 10, 50, 44, 49, 66}

	for i, v := range values {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", values)
}
