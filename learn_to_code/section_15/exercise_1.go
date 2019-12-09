package main

import "fmt"

func main() {
	fmt.Println(foo(5, 4))

	i, s := bar()
	fmt.Println(i, s)
}

func foo(x int, y int) int {
	return x + y*y
}

func bar() (int, string) {
	return 5, "hello"
}
