package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return dividend / divisor, nil
}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
}
