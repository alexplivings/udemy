package main

import "fmt"

type character struct {
	name string
	hp   int
	str  int
	def  int
}

// SimulateBattle simulates battling two characters
func SimulateBattle(c1, c2 *character) (bool, error) {
	c1.hp = c1.hp - (c2.str - c1.def)
	c2.hp = c2.hp - (c1.str - c2.def)
	return false, nil
}

func main() {
	c1 := character{
		name: "Faelvien",
		hp:   100,
		str:  15,
		def:  5,
	}

	c2 := character{
		name: "Defavaer",
		hp:   112,
		str:  11,
		def:  8,
	}

	SimulateBattle(&c1, &c2)

	fmt.Printf("%v\n\n%v\n", c1, c2)
}
