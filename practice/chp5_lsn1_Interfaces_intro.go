package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func sendMessage(msg message) (string, int) {
	msgText := msg.getMessage()
	costOfMessage := len(msgText) * 3

	return msgText, costOfMessage
}

func main() {
	bMessage := birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"}
	sReport := sendingReport{"First Report", 10}

	fmt.Println(bMessage.getMessage())
	fmt.Println(sReport.getMessage())

	fmt.Println(sendMessage(bMessage))
}
