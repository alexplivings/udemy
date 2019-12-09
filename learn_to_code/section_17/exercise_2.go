package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	isIdiot   bool
}

func changeName(p *person, firstName string, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
}

func main() {
	p1 := person{
		firstName: "Billy",
		lastName:  "Schmidt",
		age:       12,
		isIdiot:   false,
	}

	fmt.Println(p1)

	changeName(&p1, "Billington", "Schmidtington")

	fmt.Println(p1)
}
