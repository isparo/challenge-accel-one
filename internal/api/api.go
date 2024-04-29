package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/josue/challenge-accel-one/internal/api/handler"
	domaincontact "github.com/josue/challenge-accel-one/internal/domain/contact"

	"github.com/josue/challenge-accel-one/internal/infrastructure/persistency"
	"github.com/josue/challenge-accel-one/internal/shared/apiconfig"
)

type contactHandler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type apiV1 struct {
	//interface property to inject handlers
	contactHandler contactHandler
	apiConfig      *apiconfig.APIConfig
}

func newApiV1(apiConfig *apiconfig.APIConfig, contactHandler contactHandler) apiV1 {
	return apiV1{
		contactHandler: contactHandler,
		apiConfig:      apiConfig,
	}
}

func LoadAPIV1() {
	apiConfig, err := apiconfig.NewAPIConfig()
	if err != nil {
		log.Fatal("error loading API configuration: ", err.Error())
	}

	inMemoryRepo := persistency.NewInMemoryRepository()

	contactSvc := domaincontact.NewContactService(&inMemoryRepo)
	contactHandler := handler.Newcontacthandler(contactSvc)
	apiV1 := newApiV1(&apiConfig, contactHandler)

	apiV1.LoadRoutes()
}
