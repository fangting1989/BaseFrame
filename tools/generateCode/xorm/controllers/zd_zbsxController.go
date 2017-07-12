package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func Registerzd_zbsxRoutes(route *gin.Engine) {
	 zd_zbsxController := route.Group("zd_zbsxController")
	{
		zd_zbsxController.GET("/insert",zd_zbsxInsert)
		zd_zbsxController.GET("/update",zd_zbsxUpdate)
		zd_zbsxController.GET("/del",zd_zbsxDel)
		zd_zbsxController.GET("/find",zd_zbsxFind)
	}
}
/*对象查询*/
func zd_zbsxFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "zd_zbsxFind",
	})
}
/*对象插入*/
func zd_zbsxInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "zd_zbsxInsert",
	})
}
/*对象更新*/
func zd_zbsxUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "zd_zbsxUpdate",
	})
}
/*对象删除*/
func zd_zbsxDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "zd_zbsxDel",
	})}