package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
)

type ContactRequest struct {
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

// ** ** This comment is to the code evaluator **
// In this section we can add more validations
// we can add custom validations if is needed
func (cr ContactRequest) Validate() error {
	return validation.ValidateStruct(&cr,
		validation.Field(&cr.Name, validation.Required),
		validation.Field(&cr.LastName, validation.Required),
		validation.Field(&cr.Email, validation.Required),
		validation.Field(&cr.PhoneNumber, validation.Required))
}

type ContactResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
}

func NewContactResponse(
	id uuid.UUID,
	name string,
	lastName string,
	email string,
	phoneNumber string,
) ContactResponse {
	return ContactResponse{
		ID:          id,
		Name:        name,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}
