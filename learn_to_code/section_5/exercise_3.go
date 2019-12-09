package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v is %v years old, yes that is %v", y, x, z)
	fmt.Println(s)
}
