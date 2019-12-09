package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(y, "is", x, "years old, yes that is", z)
	fmt.Printf("%v is %v years old, yes that is %v", y, x, z)
}
