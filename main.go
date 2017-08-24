package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"./middleware"
	"./routers"
	"./utils"
)

func main() {
	//初始化路由
	Router := gin.Default()
	Router.Use(middleware.Jwtmiddleware())
	Router.Use(middleware.Sqlmiddleware())

	r := routers.InitRouter(Router)
	//绑定端口
	// gin.SetMode(gin.ReleaseMode)
	http.ListenAndServe(":"+utils.StartPort(), r)
}
