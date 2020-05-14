package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println("The area is", s.area())
}

func main() {
	s := square{
		10.0,
	}
	info(s)

	c := circle{
		5.0,
	}
	info(c)
}
