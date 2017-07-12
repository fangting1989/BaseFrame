package routers

import (
	"net/http"

	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() http.Handler {
	router := gin.Default()
	// var abc = controllers
	controllers.Registercjgs_zdRoutes(router)
	controllers.Registercrgy_kzRoutes(router)
	controllers.RegistergdxmRoutes(router)
	return router
}
