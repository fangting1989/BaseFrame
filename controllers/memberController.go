package controllers

import (
	"github.com/gin-gonic/gin"
)

//
func RegistermemberRoutes(route *gin.Engine) {
	memberController := route.Group("memberController")
	{
		memberController.GET("/find", memberFind)
		memberController.GET("/insert", memberInsert)
		memberController.GET("/update", memberUpdate)
		memberController.GET("/del", memberDel)
	}
}

/*对象查询*/
func memberFind(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "memberFind",
	})
}

func memberInsert(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "memberInsert",
	})
}

func memberUpdate(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "memberUpdate",
	})
}

func memberDel(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "memberDel",
	})
}
