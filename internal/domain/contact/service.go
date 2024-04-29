package contact

import uuid "github.com/satori/go.uuid"

type contactService struct {
	contactRepository ContactRepository
}

func NewContactService(contactRepo ContactRepository) contactService {
	return contactService{
		contactRepository: contactRepo,
	}
}

func (cs contactService) Create(
	name string,
	lastName string,
	email string,
	phoneNumber string,
) (*uuid.UUID, error) {

	contact := Contact{
		Name:        name,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}

	id, err := cs.contactRepository.Create(contact)
	if err != nil {
		//add logs
		return nil, err
	}

	return id, nil
}

func (cs contactService) GetByID(id uuid.UUID) (*Contact, error) {

	//more posible code here

	contact, err := cs.contactRepository.GetByID(id)

	if err != nil {
		//add logs here
		return nil, err
	}

	return contact, nil
}

func (cs contactService) Delete(id uuid.UUID) error {

	//more posible code here

	err := cs.contactRepository.Delete(id)

	if err != nil {
		//add logs here
		return err
	}

	return nil
}

func (cs contactService) Update(
	id uuid.UUID,
	name string,
	lastName string,
	email string,
	phoneNumber string,
) error {

	// more posible code here

	contact := NewContact(
		id,
		name,
		lastName,
		email,
		phoneNumber,
	)

	err := cs.contactRepository.Update(contact)
	if err != nil {
		//add logs here
		return err
	}

	return nil
}
