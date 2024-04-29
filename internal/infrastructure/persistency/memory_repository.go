package persistency

import (
	"sync"

	domaincontact "github.com/josue/challenge-accel-one/internal/domain/contact"
	"github.com/josue/challenge-accel-one/internal/shared/errorhandler"
	uuid "github.com/satori/go.uuid"
)

type inMemoryRepository struct {
	dataStorage map[uuid.UUID]domaincontact.Contact
	mtx         sync.Mutex
}

func NewInMemoryRepository() inMemoryRepository {
	return inMemoryRepository{
		dataStorage: make(map[uuid.UUID]domaincontact.Contact),
	}
}

func (r *inMemoryRepository) Create(contact domaincontact.Contact) (*uuid.UUID, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	contactID := uuid.NewV4()
	contact.ID = contactID

	r.dataStorage[contactID] = contact

	return &contactID, nil
}

func (r *inMemoryRepository) GetByID(id uuid.UUID) (*domaincontact.Contact, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	contact, ok := r.dataStorage[id]

	if !ok {
		return nil, errorhandler.NewRecordNotFoundError("contact not found", errorhandler.RecordNotFound)
	}

	return &contact, nil
}

func (r *inMemoryRepository) Delete(id uuid.UUID) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.dataStorage[id]; !ok {
		return errorhandler.NewRecordNotFoundError("contact not found", errorhandler.RecordNotFound)
	}

	delete(r.dataStorage, id)

	return nil
}

func (r *inMemoryRepository) Update(contact domaincontact.Contact) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.dataStorage[contact.ID]; !ok {
		return errorhandler.NewRecordNotFoundError("contact not found", errorhandler.RecordNotFound)
	}

	r.dataStorage[contact.ID] = contact

	return nil
}
