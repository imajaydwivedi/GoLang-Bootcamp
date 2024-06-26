package main

import "fmt"

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense) (string, float64) {

	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return fmt.Sprintf("%T\n", v), 0.0
	}
}

func main() {
	//var e expense = email{isSubscribed: true, body: "thanks for subscribing!", toAddress: "abc@gmail.com"}
	e := email{isSubscribed: true, body: "thanks for subscribing!", toAddress: "abc@gmail.com"}
	s := sms{isSubscribed: false, body: "thanks for subscribing!", toPhoneNumber: "123456789"}

	fmt.Println(getExpenseReport(e))
	fmt.Println(getExpenseReport(s))
}
