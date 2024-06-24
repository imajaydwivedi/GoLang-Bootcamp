package main

import (
	"fmt"
)

type car struct {
	make    string
	model   string
	door    int
	mileage int
}

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {
	msg2Send := messageToSend{123456789, "Thanks for details."}
	fmt.Printf("Message \"%v\" sent to %v\n", msg2Send.message, msg2Send.phoneNumber)
}
