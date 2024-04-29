package main

import (
	"github.com/josue/challenge-accel-one/internal/api"

	_ "github.com/josue/challenge-accel-one/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@title			Contact API
//	@version		1.0
//	@description	APIs to manage contacts.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	api.LoadAPIV1()
}
