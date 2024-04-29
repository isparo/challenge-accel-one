package contact

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/josue/challenge-accel-one/internal/shared/errorhandler"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockrep := NewMockContactRepository(ctrl)
	contactSvc := NewContactService(mockrep)

	testUUID := uuid.NewV4()
	testName := "testName"
	testLastName := "testLastName"
	testPhone := "123456"
	testEmail := "theemail@some.com"

	testError := errorhandler.NewRecordNotFoundError("contact not found", errorhandler.RecordNotFound)

	tt := []struct {
		name       string
		err        error
		wantsErr   bool
		contactID  uuid.UUID
		resContact *Contact
	}{
		{
			name:      "success",
			err:       nil,
			wantsErr:  false,
			contactID: testUUID,
			resContact: &Contact{
				ID:          testUUID,
				Name:        testName,
				LastName:    testLastName,
				Email:       testEmail,
				PhoneNumber: testPhone,
			},
		},
		{
			name:       "contact not found",
			err:        testError,
			wantsErr:   true,
			contactID:  testUUID,
			resContact: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			mockrep.EXPECT().
				GetByID(gomock.Any()).
				Times(1).
				Return(tc.resContact, tc.err)

			contact, err := contactSvc.GetByID(tc.contactID)

			if tc.wantsErr {
				assert.NotNil(t, err)
				assert.Nil(t, contact)
				assert.Contains(t, err.Error(), "record not found")
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, contact)
				assert.Equal(t, contact.ID, tc.contactID)
				assert.Equal(t, contact.Name, tc.resContact.Name)
				assert.Equal(t, contact.LastName, tc.resContact.LastName)
				assert.Equal(t, contact.Email, tc.resContact.Email)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrep := NewMockContactRepository(ctrl)

	contactSvc := NewContactService(mockrep)

	testName := "testName"
	testLastName := "testLastName"
	testPhone := "123456"
	testEmail := "theemail@some.com"
	testErr := "testError"

	testUUID := uuid.NewV4()

	tt := []struct {
		name       string
		err        error
		wantsErr   bool
		obj        Contact
		expectedID *uuid.UUID
	}{
		{
			name:       "success",
			err:        nil,
			wantsErr:   false,
			expectedID: &testUUID,
			obj: Contact{
				Name:        testName,
				LastName:    testLastName,
				Email:       testEmail,
				PhoneNumber: testPhone,
			},
		},
		{
			name:       "failure",
			err:        errors.New(testErr),
			wantsErr:   true,
			expectedID: nil,
			obj: Contact{
				Name:        testName,
				LastName:    testLastName,
				Email:       testEmail,
				PhoneNumber: testPhone,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			mockrep.EXPECT().
				Create(gomock.Any()).
				Times(1).
				Return(tc.expectedID, tc.err)

			id, err := contactSvc.Create(tc.obj.Name, tc.obj.LastName,
				tc.obj.Email, tc.obj.PhoneNumber)

			if tc.wantsErr {
				assert.NotNil(t, err)
				assert.Nil(t, id)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, id)
			}

		})
	}
}
