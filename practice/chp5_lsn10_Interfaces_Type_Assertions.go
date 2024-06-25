package main

import (
	"fmt"
	"log"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	height int
	width  int
}

func (s square) area() float64 {
	return float64(s.height) * float64(s.width)
}

func main() {

	//var s shape = circle{radius: 5}
	var s shape = square{height: 5, width: 10}

	// if ok is true, then c is a circle type object
	c, ok := s.(circle)
	if !ok {
		// log an error if s isn't a circle
		log.Fatal("s is not a circle")
	} else {
		fmt.Println("s is a cicle type")
		fmt.Printf("Circle area = %.2f\n", c.area())
	}

}
