package main

import (
	"fmt"
)

type car struct {
	make    string
	model   string
	doors   int
	mileage int

	wheel struct {
		radius   int
		material string
	}
}

func main() {

	myCar := struct {
		make  string
		model string
	}{
		make:  "tesla",
		model: "model 3",
	}

	fmt.Printf("%v car of %v\n", myCar.make, myCar.model)

	myCar2 := car{}
	if myCar2.make == "" {
		fmt.Println("struct variable not assigned.\n")
	}
}
