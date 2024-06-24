package main

import "fmt"

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType membershipType) User {
	user := User{}

	user.Name = name
	user.Type = membershipType
	if user.Type == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}

	return user
}

func main() {
	newUser := newUser("Ajay", TypePremium)

	fmt.Printf("%s is a %s user with %d messages.\n", newUser.Name, newUser.Type, newUser.MessageCharLimit)
}
