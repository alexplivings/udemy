package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello I am", p.first, p.last, "I am", p.age, "years old, suckle on my teets")
}

func main() {
	p1 := person{
		first: "Alex",
		last:  "Livingston",
		age:   27,
	}

	p1.speak()
}
