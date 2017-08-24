package controllers

import (
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", Registercjgs_zdRoutes)
}

func Registercjgs_zdRoutes(route *gin.Engine) {
	cjgs_zdController := route.Group("cjgs_zdController")
	{
		cjgs_zdController.GET("/insert", cjgs_zdInsert)
		cjgs_zdController.GET("/update", cjgs_zdUpdate)
		cjgs_zdController.GET("/del", cjgs_zdDel)
		cjgs_zdController.GET("/find", cjgs_zdFind)
	}
}

/*对象查询*/
func cjgs_zdFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "cjgs_zdFind",
	})
}

/*对象插入*/
func cjgs_zdInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "cjgs_zdInsert",
	})
}

/*对象更新*/
func cjgs_zdUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "cjgs_zdUpdate",
	})
}

/*对象删除*/
func cjgs_zdDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "cjgs_zdDel",
	})
}
