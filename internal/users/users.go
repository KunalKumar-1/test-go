package users

import (
	"fmt"
	"net/mail"
)

type User struct {
	FirstName string
	LastName  string
	Email     mail.Address
}

type Manager struct {
	users []User
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) AddUser(firstName string, lastName string, email string) error {
	if firstName == "" {
		return fmt.Errorf("invalid first name: %q", firstName)
	}
	if lastName == "" {
		return fmt.Errorf("invalid last name:%q", lastName)
	}

	parsedAddress, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email: %s", email)
	}

	newUser := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     *parsedAddress,
	}

	m.users = append(m.users, newUser)

	return nil
}
