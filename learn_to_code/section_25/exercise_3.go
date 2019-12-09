package main

import "fmt"

type customErr struct {
	errorNo int
}

func (err customErr) Error() string {
	return fmt.Sprintf("Custom error time! %v", err.errorNo)
}

func main() {

	err := customErr{
		errorNo: 13,
	}

	foo(err)

}

func foo(e error) {
	fmt.Println(e)
}
