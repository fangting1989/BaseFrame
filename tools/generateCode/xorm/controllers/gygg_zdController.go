package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func Registergygg_zdRoutes(route *gin.Engine) {
	 gygg_zdController := route.Group("gygg_zdController")
	{
		gygg_zdController.GET("/insert",gygg_zdInsert)
		gygg_zdController.GET("/update",gygg_zdUpdate)
		gygg_zdController.GET("/del",gygg_zdDel)
		gygg_zdController.GET("/find",gygg_zdFind)
	}
}
/*对象查询*/
func gygg_zdFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "gygg_zdFind",
	})
}
/*对象插入*/
func gygg_zdInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "gygg_zdInsert",
	})
}
/*对象更新*/
func gygg_zdUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "gygg_zdUpdate",
	})
}
/*对象删除*/
func gygg_zdDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "gygg_zdDel",
	})}