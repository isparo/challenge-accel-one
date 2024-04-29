package contact

import (
	uuid "github.com/satori/go.uuid"
)

type Contact struct {
	ID          uuid.UUID
	Name        string
	LastName    string
	Email       string
	PhoneNumber string
}

func NewContact(
	id uuid.UUID,
	name string,
	lastName string,
	email string,
	phoneNumber string,
) Contact {
	return Contact{
		ID:          id,
		Name:        name,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}
