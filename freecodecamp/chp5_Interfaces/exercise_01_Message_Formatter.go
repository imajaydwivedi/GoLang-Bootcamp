package main

import (
	"fmt"
)

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

func (pt PlainText) Format() string {
	return fmt.Sprintf("%s", pt.message)
}

type Bold struct {
	message string
}

func (bd Bold) Format() string {
	return fmt.Sprintf("**%s**", bd.message)
}

type Code struct {
	message string
}

func (cd Code) Format() string {
	return fmt.Sprintf("`%s`", cd.message)
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}

func main() {
	var code Formatter = Code{message: "<bold>Hello World!</bold>"}

	fmt.Println(SendMessage(code))
}
