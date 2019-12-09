package main

import "fmt"

func main() {

	year := 1992
	for year <= 2019 {
		fmt.Println(year)
		year++
	}

	year = 1992
	for {
		fmt.Println(year)
		if year == 2019 {
			break
		}
		year++
	}

}
