package main

import "fmt"

type membershipType string

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType membershipType) User {
	membership := Membership{Type: membershipType}
	if membershipType == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membershipType == TypePremium {
		membership.MessageCharLimit = 1000
	}
	return User{Name: name, Membership: membership}
}

func (user User) SendMessage(message string, messageLength int) (string, bool) {
	if user.MessageCharLimit >= messageLength {
		return message, true
	} else {
		return "", false
	}
}

func main() {
	user := newUser("Ajay", TypePremium)

	fmt.Println(user)
	fmt.Println(user.SendMessage("Happy Birthday!", 2000))
}
