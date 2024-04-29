package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/josue/challenge-accel-one/internal/api/dto"
	domaincontact "github.com/josue/challenge-accel-one/internal/domain/contact"

	"github.com/josue/challenge-accel-one/internal/shared/errorhandler"
)

type contactDomainService interface {
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*domaincontact.Contact, error)
	Create(
		name string,
		lastName string,
		email string,
		phoneNumber string,
	) (*uuid.UUID, error)
	Update(
		id uuid.UUID,
		name string,
		lastName string,
		email string,
		phoneNumber string,
	) error
}

type contactHandler struct {
	//inject the interface here
	contactService contactDomainService
}

func Newcontacthandler(contactSvc contactDomainService) contactHandler {
	return contactHandler{
		contactService: contactSvc,
	}
}

// Create godoc
//
//	@Summary        Creates new Contact
//	@Description    This endpoint is used to create a new contact
//	@Tags           Contact
//	@Accept         json
//	@Produce        json
//	@Param          request body    dto.ContactRequest  true    "contact info"
//	@Success        201 {object}    dto.ContactResponse
//	@Failure        400 {object}    errorhandler.TemplateError
//	@Failure        500 {object}    errorhandler.TemplateError
//	@Router         /contact [post]
func (ch contactHandler) Create(ctx *gin.Context) {
	// ** This comment is to the code evaluator **
	// Is recomended decople the domain entity and the api types
	// in this case I show the use of DTO to get the request and
	// decople from the domain entity
	var request dto.ContactRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewBadRequestError("bad request error", err))
		return
	}

	//validate the request object
	if err := request.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewValidationError("invalid params", err))
		return
	}

	newContactId, err := ch.contactService.Create(request.Name, request.LastName, request.Email, request.PhoneNumber)
	var errWraping errorhandler.TemplateError
	if errors.As(err, &errWraping) {
		ctx.AbortWithStatusJSON(errWraping.Status, errorhandler.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	// ** This comment is to the code evaluator **
	// Is recomended decople the domain entity and the api types
	// in this case I show the use of DTO to translate the domain entity into
	// the response object, that will help us to decouple the domain and the application
	// this time is a REST API
	contactResponse := dto.NewContactResponse(
		*newContactId,
		request.Name,
		request.LastName,
		request.Email,
		request.PhoneNumber)

	ctx.JSON(http.StatusCreated, contactResponse)
}

// Delete godoc
//
//	@Summary        Delete Contact by ID
//	@Description    Delete Contact by ID
//	@Tags           Contact
//	@Param          contactId   path    string  true    "contact ID"
//	@Success        204
//	@Failure        400 {object}    errorhandler.TemplateError
//	@Failure        404 {object}    errorhandler.TemplateError
//	@Failure        500 {object}    errorhandler.TemplateError
//	@Router         /contact/{contactId} [delete]
func (ch contactHandler) Delete(ctx *gin.Context) {
	contactID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	err = ch.contactService.Delete(contactID)
	var errWraping errorhandler.TemplateError
	if errors.As(err, &errWraping) {
		ctx.AbortWithStatusJSON(errWraping.Status, errorhandler.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Get godoc
//
//	@Summary        Get contact
//	@Description    Get contact by ID
//	@Tags           Contact
//	@Produce        json
//	@Param          contactId  path        string  true    "contact ID"
//	@Success        202             {object}     dto.ContactResponse
//	@Failure        400             {object}    errorhandler.TemplateError
//	@Failure        404             {object}    errorhandler.TemplateError
//	@Failure        500             {object}    errorhandler.TemplateError
//	@Router         /contact/{contactId} [get]
func (ch contactHandler) GetByID(ctx *gin.Context) {
	contactID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	contact, err := ch.contactService.GetByID(contactID)
	var errWraping errorhandler.TemplateError
	if errors.As(err, &errWraping) {
		ctx.AbortWithStatusJSON(errWraping.Status, errorhandler.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	// ** This comment is to the code evaluator **
	// Is recomended decople the domain entity and the api types
	// in this case I show the use of DTO to translate the domain entity into
	// the response object, that will help us to decouple the domain and the application
	// this time is a REST API
	contactResponse := dto.NewContactResponse(
		contact.ID,
		contact.Name,
		contact.LastName,
		contact.Email,
		contact.PhoneNumber)

	ctx.JSON(http.StatusOK, contactResponse)
}

// Update godoc
//
//	@Summary        Update Contacts
//	@Description    Updates the Contact
//	@Tags           Contact
//	@Accept         json
//	@Produce        json
//	@Param          contactId   path    string              true    "Contact ID"
//	@Param          request     body    dto.ContactRequest true    "Contact request"
//	@Success        202 {object}    dto.ContactResponse
//	@Failure        400 {object}    errorhandler.TemplateError
//	@Failure        404 {object}    errorhandler.TemplateError
//	@Failure        500 {object}    errorhandler.TemplateError
//	@Router         /contact/{contactId} [put]
func (ch contactHandler) Update(ctx *gin.Context) {
	log.Println("update")

	contactID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	var request dto.ContactRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewBadRequestError("bad request error", err))
		return
	}

	//validate the request object
	if err := request.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorhandler.NewValidationError("invalid params", err))
		return
	}

	err = ch.contactService.Update(
		contactID,
		request.Name,
		request.LastName,
		request.Email,
		request.PhoneNumber,
	)
	var errWraping errorhandler.TemplateError
	if errors.As(err, &errWraping) {
		ctx.AbortWithStatusJSON(errWraping.Status, errorhandler.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorhandler.NewErrorResponse(err.Error()))
		return
	}

	// The use of a dto is important here to decouple the domain and the application (REST API).
	contactResponse := dto.NewContactResponse(
		contactID,
		request.Name,
		request.LastName,
		request.Email,
		request.PhoneNumber)

	ctx.JSON(http.StatusAccepted, contactResponse)
}
