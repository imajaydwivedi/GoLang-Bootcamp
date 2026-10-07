package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	costs := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

	return messages, costs
}

func main() {
	messages, costs := getMessageWithRetries("hello", "world", "!")

	fmt.Printf("'%s' sent at cost of %v\n", messages[0], costs[0])  
	fmt.Printf("'%s' sent at cost of %v\n", messages[1], costs[1])
	fmt.Printf("'%s' sent at cost of %v\n", messages[2], costs[2])
}
 