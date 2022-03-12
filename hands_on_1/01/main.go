package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * (math.Pi * math.Pi)
}

type shape interface {
	area() float64
}

func info(s shape) {

	fmt.Printf("%v\n", s.area())

}

func main() {

	sq1 := square{
		4,
	}

	c1 := circle{
		3,
	}

	info(sq1)
	info(c1)
}
