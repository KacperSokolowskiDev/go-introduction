package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

// struct embedding
type Admin struct {
	email string
	password string
	User //anonymous embedding, allows using methods directly through Admin
}

//participant user is a receiver argument of data from the user struct, they are linked now
func (participant *User) OutputUserDetails() {
	//technically correct way is (*participant).firstName as we would need to dereference it but Go also allows the below
	fmt.Println(participant.firstName, participant.lastName, participant.birthDate)
}

func (participant *User) ClearUserName() {
	participant.firstName = ""
	participant.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}

func New(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}