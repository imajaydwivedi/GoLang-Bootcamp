package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

type authorizationInfo struct {
	username string
	password string
}

func (ai authorizationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %v:%v", ai.username, ai.password)
}

func main() {

	var r = rect{
		width:  5,
		height: 10,
	}

	fmt.Printf("area = %v\n", r.area())

	authDetails := authorizationInfo{
		username: "adwivedi",
		password: "12345678",
	}

	fmt.Println(authDetails.getBasicAuth())
}
