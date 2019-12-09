package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	myCar := sedan{
		vehicle: vehicle{doors: 4, color: "white"},
		luxury:  true,
	}

	otherCar := truck{
		vehicle:   vehicle{doors: 4, color: "blue"},
		fourWheel: true,
	}

	fmt.Println(myCar)
	fmt.Println(otherCar)

	fmt.Println(myCar.doors)
}
