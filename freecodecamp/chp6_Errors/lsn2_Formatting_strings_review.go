package main

import (
	"fmt"
	"log"
)

//https://pkg.go.dev/log

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs %.2f to be sent to '%v' can not be sent", cost, recipient)
}

func main() {
	log.Fatal(getSMSErrorString(10.0, "ajay@gmail.com"))
}
 