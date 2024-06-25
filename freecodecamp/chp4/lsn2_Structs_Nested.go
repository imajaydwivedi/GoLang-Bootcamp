package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {

	if (mToSend.sender.name == "") || (mToSend.sender.number == 0) {
		return false
	}

	if (mToSend.recipient.name == "") || (mToSend.recipient.number == 0) {
		return false
	}

	return true
}

func main() {
	//msg2Send := messageToSend{123456789, "Thanks for details."}
	msg2Send := messageToSend{}

	var sendMsg bool
	sendMsg = canSendMessage(msg2Send)

	if sendMsg {
		fmt.Println("Sender or Recipient details correct.\n")
	} else {
		fmt.Println("Sender or Recipient details missing.\n")
	}

}
