package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type fullTime struct {
	name   string
	salary int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hoursPerYear * c.hourlyPay
}

func main() {

	var ft employee = fullTime{
		name:   "Ajay",
		salary: 500000,
	}

	var c employee = contractor{
		name:         "Gaurav",
		hourlyPay:    10,
		hoursPerYear: 1000,
	}

	fmt.Println(ft)
	fmt.Printf("%s is an employee with %d salary.\n", ft.getName(), ft.getSalary())

	fmt.Println(c)
	fmt.Printf("%s is a contractor with %d salary.\n", c.getName(), c.getSalary())
}
