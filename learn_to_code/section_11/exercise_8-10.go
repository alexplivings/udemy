package main

import "fmt"

func main() {
	faves := map[string][]string{}

	faves["livingston"] = []string{"Tumbling", "Big Weed", "Computers"}
	faves["bond"] = []string{"Silenced PP7", "Tuxedos", "Silver cars"}

	faves["lewis"] = []string{"Journalism", "Funny videos", "Whisky"}

	fmt.Println(faves)

	for k, v := range faves {
		fmt.Println(k, "-", v)
	}

	fmt.Println()
	delete(faves, "lewis")

	for k, v := range faves {
		fmt.Println(k, "-", v)
	}
}
