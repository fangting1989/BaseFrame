package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegistergdxmRoutes(route *gin.Engine) {
	gdxmController := route.Group("gdxmController")
	{
		gdxmController.POST("/insert", gdxmInsert)
		gdxmController.GET("/update", gdxmUpdate)
		gdxmController.GET("/del", gdxmDel)
		gdxmController.GET("/find", gdxmFind)
	}
}

/*对象查询*/
func gdxmFind(c *gin.Context) {
	var pageobject models.Pageobject
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		if c.Bind(&pageobject) == nil {
			c.JSON(http.StatusOK, pageobject)
		} else {
			c.JSON(http.StatusOK, gdxm)
		}

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "unauthorized"})
	}
}

//find
// c.JSON(200, gin.H{
// 	"message": "gdxmFind",
// })

/*对象插入*/
func gdxmInsert(c *gin.Context) {
	//insert
	var gdxm models.Gdxm
	// if c.Bind(&gdxm) == nil { BindJSON
	if c.BindJSON(&gdxm) == nil {
		c.JSON(http.StatusOK, gdxm)

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "unauthorized"})
	}
}

/*对象更新*/
func gdxmUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "gdxmUpdate",
	})
}

/*对象删除*/
func gdxmDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "gdxmDel",
	})
}
