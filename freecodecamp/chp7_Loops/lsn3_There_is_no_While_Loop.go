package main

import "fmt"

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies < float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies >= float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func main() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")

	fmt.Println("******************************************")

	var costMultiplier float64 = 1.1
	var maxCostInPennies int = 5
	fmt.Printf("costMultiplier: %v \nmaxCostInPennies: %v \nTotal message: %v\n\n\n", costMultiplier, maxCostInPennies, getMaxMessagesToSend(costMultiplier, maxCostInPennies))

	costMultiplier = 1.3
	maxCostInPennies = 10
	fmt.Printf("costMultiplier: %v \nmaxCostInPennies: %v \nTotal message: %v\n\n\n", costMultiplier, maxCostInPennies, getMaxMessagesToSend(costMultiplier, maxCostInPennies))

	costMultiplier = 1.35
	maxCostInPennies = 25
	fmt.Printf("costMultiplier: %v \nmaxCostInPennies: %v \nTotal message: %v\n\n\n", costMultiplier, maxCostInPennies, getMaxMessagesToSend(costMultiplier, maxCostInPennies))
}
