package main

import (
	"fmt"
	"testing"
)

func TestSimulateBattle(t *testing.T) {
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

	res, err := SimulateBattle(&c1, &c2)
	if err != nil {
		fmt.Println("err:", err)
	}

	if !res {
		t.Errorf("Got false, want true.")
	}
}
