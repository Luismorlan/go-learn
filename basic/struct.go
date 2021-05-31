package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Trever",
		LastName:  "Sawler",
	}

	log.Println(user)
}

// Different scope.
func CapitalCanBeSeenOutside()      {}
func lowerCaseCannotBeSeenOutside() {}
