package main

import (
	"fmt"

	"github.com/alexplivings/udemy/learn_to_code/section_29/exercise_2/quote"
	"github.com/alexplivings/udemy/learn_to_code/section_29/exercise_2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
