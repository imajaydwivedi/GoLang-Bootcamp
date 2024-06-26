package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Main called.")

	rs, _ := divide(10, 2)
	fmt.Printf("%v", rs)

}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

func getInsuranceAmount(status insuranceStatus) int {

	amount := 0
	if !status.hasInsurance() {
		amount = 1
	} else {
		if status.isTotaled() {
			amount = 10000
		} else {
			if status.isDented() {
				amount = 160
				if status.isBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}

func getInsuranceAmount_With_Gaurd_Clauses(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}
