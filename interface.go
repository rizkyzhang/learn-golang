package main

import (
	"fmt"
	"math"
)

func learnInterface() {
	var shape calc
	shape = square{5}

	fmt.Println("Square area:", shape.area())
	fmt.Println("Square circumference:", shape.circumference())

	shape = circle{10}

	fmt.Println("Circle area:", shape.area())
	fmt.Println("Circle circumference:", shape.circumference())
	fmt.Println("Circle radius", shape.(circle).radius())

}

type calc interface {
	area() float64
	circumference() float64
}

// Circle
type circle struct {
	diameter float64
}

func (c circle) radius() float64 {
	return c.diameter / 2
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) circumference() float64 {
	return math.Pi * c.diameter
}

// Square
type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s square) circumference() float64 {
	return s.side * 4
}

