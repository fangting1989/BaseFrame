package main

import (
	"net/http"

	"./routers"
)

func main() {
	//初始化路由
	router := routers.InitRouter()
	//绑定端口
	http.ListenAndServe(":8000", router)
}
