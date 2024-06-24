package main

import "fmt"

type car struct {
	make  string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a "truck" now also additionally contains all of the fields of the car struct
	car
	bedsize int
}

type user struct {
	name   string
	number int
}

type sender struct {
	user
	rateLimit int
}

func main() {

	lanesTruck := truck{
		bedsize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedsize)

	//embedded fields promoted to the top-level
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)

	mySender := sender{
		rateLimit: 10,
		user: user{
			name:   "Ajay",
			number: 123456789,
		},
	}

	fmt.Printf("%v is allowed to send %v messages using number %v\n", mySender.name, mySender.rateLimit, mySender.number)
}
