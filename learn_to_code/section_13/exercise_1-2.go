package main

import "fmt"

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {
	alex := person{
		firstName:  "Alex",
		lastName:   "Livingston",
		favFlavors: []string{"Coffee", "Gingerbread", "Chocolate"},
	}
	carl := person{
		firstName:  "Carl",
		lastName:   "Livingston",
		favFlavors: []string{"Mint chocolate chip", "Strawberry", "Rocky road"},
	}

	fmt.Println(alex)
	fmt.Println(carl)

	people := map[string]person{}

	people["Alex"] = alex
	people["Carl"] = carl

	for _, v := range people {
		fmt.Println(v)
	}
}
