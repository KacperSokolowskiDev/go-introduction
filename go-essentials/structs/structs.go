package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (DD/MM/YYYY): ")

	var appUser User
	//can use shorter notation (below) when elements from struct are in correct order
	appUser = User {
		userFirstName,
		userLastName,
		userBirthDate,
		time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(participant *User) {
	//technically correct way is (*participant).firstName as we would need to dereference it but Go also allows the below
	fmt.Println(participant.firstName, participant.lastName, participant.birthDate, participant.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}