package main

import "fmt"

func main() {
	defer closeYourMom()

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))

	fmt.Println(summmm(numbers))
}

func closeYourMom() {
	fmt.Println("Your mom is now closed.")
}

func sum(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func summmm(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
