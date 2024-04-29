package contact

import (
	uuid "github.com/satori/go.uuid"
)

type ContactRepository interface {
	Create(contact Contact) (*uuid.UUID, error)
	GetByID(id uuid.UUID) (*Contact, error)
	Delete(id uuid.UUID) error
	Update(contact Contact) error
}
