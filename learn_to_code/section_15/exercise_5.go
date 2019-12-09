package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
	width  int
}

type circle struct {
	radius int
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	area := s.area()
	return area
}

func (c circle) area() float64 {
	return 2 * float64(c.radius) * math.Pi
}

func (s square) area() float64 {
	return float64(s.length) * float64(s.width)
}

func main() {
	s1 := square{
		length: 10,
		width:  10,
	}

	c1 := circle{
		radius: 5,
	}

	s1area := info(s1)
	c1area := info(c1)

	fmt.Println(s1area)
	fmt.Println(c1area)
}
