package routers

import (
	"net/http"

	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() http.Handler {
	router := gin.Default()
	// var abc = controllers
	controllers.RegistermemberRoutes(router)
	controllers.RegistermembillRoutes(router)
	return router
}
