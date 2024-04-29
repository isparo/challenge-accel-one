package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func (api apiV1) LoadRoutes() {
	r := gin.Default()
	r.GET("/ping", pong)

	apiPrefix := fmt.Sprintf("/api/%s", api.apiConfig.Version)

	v1 := r.Group(apiPrefix)
	{
		contact := v1.Group(api.apiConfig.ContactGroup)
		{
			contact.GET("ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"msg": "pong"})
			})
			contact.GET(":id", api.contactHandler.GetByID)
			contact.POST("", api.contactHandler.Create)
			contact.PUT(":id", api.contactHandler.Update)
			contact.DELETE(":id", api.contactHandler.Delete)
		}
	}

	// load swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(api.apiConfig.Host + ":" + api.apiConfig.Port)
	if err != nil {
		log.Println("can not start service")
	}

}
