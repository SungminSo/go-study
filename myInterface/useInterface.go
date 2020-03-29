package main

import (
	myInterface2 "awesomeProject/src/myInterface"
	"fmt"
	"math"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (s circle) Area() float64 {
	return s.R * s.R * math.Pi
}

func (s circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

func CalCulate(x myInterface2.Shape) {
	_, ok := x.(circle)

	if ok {
		fmt.Println("Is a circle!")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("Is a sqaure:", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	CalCulate(x)
	y := circle{R: 5}
	CalCulate(y)
}