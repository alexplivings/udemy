package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi, I am", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Alex",
		last:  "Livingston",
		age:   27,
	}

	saySomething(p1)

	saySomething(&p1)
}
