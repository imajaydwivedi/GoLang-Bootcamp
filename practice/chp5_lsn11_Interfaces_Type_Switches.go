package main

import (
	"fmt"
)

type circle struct {
	radius int
}

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	var age int = 32
	c := circle{radius: 5}

	printNumericValue(age)

	printNumericValue("1")

	printNumericValue(struct{}{})

	printNumericValue(c)
}
