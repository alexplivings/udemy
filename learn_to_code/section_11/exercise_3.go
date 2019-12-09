package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	x1 := x[0:5]
	fmt.Println(x1)

	x2 := x[5:]
	fmt.Println(x2)

	x3 := x[2:7]
	fmt.Println(x3)

	x4 := x[1:6]
	fmt.Println(x4)
}

/*
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/
