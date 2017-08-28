package routers

import (
	"net/http"

	"../controllers"
	"../utils"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter(r *gin.Engine) http.Handler {
	controllers.InitControllers()
	utils.RouterBus.Publish("router:register", r)
	return r
}
