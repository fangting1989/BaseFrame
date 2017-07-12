package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func RegistertdjbxxRoutes(route *gin.Engine) {
	 tdjbxxController := route.Group("tdjbxxController")
	{
		tdjbxxController.GET("/insert",tdjbxxInsert)
		tdjbxxController.GET("/update",tdjbxxUpdate)
		tdjbxxController.GET("/del",tdjbxxDel)
		tdjbxxController.GET("/find",tdjbxxFind)
	}
}
/*对象查询*/
func tdjbxxFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "tdjbxxFind",
	})
}
/*对象插入*/
func tdjbxxInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "tdjbxxInsert",
	})
}
/*对象更新*/
func tdjbxxUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "tdjbxxUpdate",
	})
}
/*对象删除*/
func tdjbxxDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "tdjbxxDel",
	})}