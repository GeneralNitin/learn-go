package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{22.713}
	r1 := rectangle{3, 6}

	shapes := []shape{c1, r1}

	for _, s := range shapes {
		//fmt.Println("Area:", s.area())
		fmt.Println("Area:", getArea(s))
	}
}
