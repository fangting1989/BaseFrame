package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func RegistergyggRoutes(route *gin.Engine) {
	 gyggController := route.Group("gyggController")
	{
		gyggController.GET("/insert",gyggInsert)
		gyggController.GET("/update",gyggUpdate)
		gyggController.GET("/del",gyggDel)
		gyggController.GET("/find",gyggFind)
	}
}
/*对象查询*/
func gyggFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "gyggFind",
	})
}
/*对象插入*/
func gyggInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "gyggInsert",
	})
}
/*对象更新*/
func gyggUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "gyggUpdate",
	})
}
/*对象删除*/
func gyggDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "gyggDel",
	})}