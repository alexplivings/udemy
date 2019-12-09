package main

import "fmt"

func createFunc() func() {
	funk := func() {
		fmt.Println("This func was created in a func, HOLY FUNC")
	}

	return funk
}

func callbackFunc(eatCheese func()) {
	fmt.Printf("Getting ready to eat cheese, let's see what kind\n\t")

	eatCheese()
}

func main() {
	// Anonymous func
	func() {
		fmt.Println("look at this anonymous function")
	}()

	// Anon func assigned to variable
	anon := func() {
		fmt.Println("this one is assigned to a variable, WOW!")
	}
	anon()

	// Func returned by a func
	anonFunk := createFunc()
	anonFunk()

	// "callback" passing a func into a func and calling it
	eatBrie := func() {
		fmt.Println("I'm eating brie, it's my favorite cheese")
	}
	eatMozz := func() {
		fmt.Println("I'm eating mozzarelala, lalala cheese yay")
	}
	callbackFunc(eatBrie)
	callbackFunc(eatMozz)

}
