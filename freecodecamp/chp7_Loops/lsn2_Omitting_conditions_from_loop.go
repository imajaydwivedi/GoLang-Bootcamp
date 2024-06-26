package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	totalCost := 0.0
	//Each message costs 1.0, plus an additional fee. The fee structure is:
	//1st message: 1.0 + 0.00
	//2nd message: 1.0 + 0.01
	//3rd message: 1.0 + 0.02
	//4th message: 1.0 + 0.03
	//...

	for i := 0; ; i++ {
		totalCost += 1.0 + (float64(i) * 0.01)
		if totalCost > thresh {
			return i
		}
	}
}

func bulkSend(numMessages int) float64 {
	var cost float64 = 0.0

	for i := 0; i < numMessages; i++ {
		cost += 1.0 + (float64(i) * 0.01)
	}

	return cost
}

func main() {
	fmt.Printf("Cost of 10 messages = %.2f\n", bulkSend(10))
	fmt.Printf("Cost of 50 messages = %.2f\n", bulkSend(50))
	fmt.Printf("Cost of 100 messages = %.2f\n", bulkSend(100))

	fmt.Println("*********************************")

	fmt.Printf("No of messages for 10  = %v\n", maxMessages(10))
	fmt.Printf("No of messages for 50 = %v\n", maxMessages(50))
	fmt.Printf("No of messages for 150 = %v\n", maxMessages(150))
}
