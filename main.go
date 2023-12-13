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
		FirstName: "James",
		LastName:  "Bond",
		PhoneNumber: "444-666-888",
	}

	log.Println(user.FirstName, user.LastName, "birthdate:", user.BirthDate,user.PhoneNumber)
}
