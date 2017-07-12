package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func RegisterzddxzqRoutes(route *gin.Engine) {
	 zddxzqController := route.Group("zddxzqController")
	{
		zddxzqController.GET("/insert",zddxzqInsert)
		zddxzqController.GET("/update",zddxzqUpdate)
		zddxzqController.GET("/del",zddxzqDel)
		zddxzqController.GET("/find",zddxzqFind)
	}
}
/*对象查询*/
func zddxzqFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "zddxzqFind",
	})
}
/*对象插入*/
func zddxzqInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "zddxzqInsert",
	})
}
/*对象更新*/
func zddxzqUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "zddxzqUpdate",
	})
}
/*对象删除*/
func zddxzqDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "zddxzqDel",
	})}