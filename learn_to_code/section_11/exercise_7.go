package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	z := [][]string{x, y}
	fmt.Println(len(z))
	fmt.Println(cap(z))

	for _, v := range z {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}

}
