package controllers

import (
	"github.com/gin-gonic/gin"
)

//
func RegistermembillRoutes(route *gin.Engine) {
	membillController := route.Group("membillController")
	{
		membillController.GET("/insert", membillInsert)
		membillController.GET("/update", membillUpdate)
		membillController.GET("/del", membillDel)
		membillController.GET("/find", membillFind)
	}
}

/*对象查询*/
func membillFind(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "membillFind",
	})
}

func membillInsert(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "membillInsert",
	})
}

func membillUpdate(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "updateInsert",
	})
}

func membillDel(c *gin.Context) {
	//user login
	c.JSON(200, gin.H{
		"message": "membillDel",
	})
}
