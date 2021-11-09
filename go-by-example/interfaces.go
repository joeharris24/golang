package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures

type geometry interface {
	area() float64
	//perim() float64
}

// Implementing this interface on rect and circle types

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement interface, need to implement all the methods in the interface
// Implementation of geometry on rect

func (r rect) area() float64 {
	return r.width * r.height
}

// Implementation of geometry on circle

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
