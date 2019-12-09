package main

import "fmt"

func main() {
	values := [5]int{1, 2, 3, 4, 5}

	fmt.Println(values)

	values[3] = 7

	fmt.Println(values)

	for i, v := range values {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", values)
}
