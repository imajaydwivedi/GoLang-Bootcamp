package main

import (
	"fmt"
)

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costToCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	costToSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return costToCustomer + costToSpouse, nil
}

func main() {
	fmt.Println(sendSMSToCouple("Reminder for Anniversary!", "Happy Anniversary Wife! \nI don't know what else to say.\nWas thinking if you are interested for a family outing?"))
	fmt.Println(sendSMSToCouple("Reminder for Anniversary!", "Happy Anniversary Wife!"))
}
