package routers

import (
	"net/http"

	"../controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter(r *gin.Engine) http.Handler {
	controllers.InitControllers()
	// utils.RouterBus.Publish("router:register", r)
	controllers.RegistergdxmRoutes(r)
	// r.ServeHTTP(w, req){
	// 	fmt.Println("--request")
	// }

	return r
}
