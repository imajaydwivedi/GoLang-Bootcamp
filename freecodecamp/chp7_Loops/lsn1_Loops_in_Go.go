package main

import (
	"fmt"
)

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
}
