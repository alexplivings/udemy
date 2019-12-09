package main

import "fmt"

const (
	thisYear         int = 2019 + iota
	nextYear         int = thisYear + iota
	nextNextYear     int = thisYear + iota
	nextNextNextYear int = thisYear + iota
)

func main() {
	x := 31

	fmt.Printf("%d\t\t%b\t\t%X\n", x, x, x)

	x = x << 1

	fmt.Printf("%d\t\t%b\t\t%X\n\n", x, x, x)

	y := ` holy boi		this looks
like a funnn
	
	
	
	thing to do`

	fmt.Println(y)

	fmt.Println(thisYear, nextYear, nextNextYear, nextNextNextYear)
}
