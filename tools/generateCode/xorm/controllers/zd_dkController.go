package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func Registerzd_dkRoutes(route *gin.Engine) {
	 zd_dkController := route.Group("zd_dkController")
	{
		zd_dkController.GET("/insert",zd_dkInsert)
		zd_dkController.GET("/update",zd_dkUpdate)
		zd_dkController.GET("/del",zd_dkDel)
		zd_dkController.GET("/find",zd_dkFind)
	}
}
/*对象查询*/
func zd_dkFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "zd_dkFind",
	})
}
/*对象插入*/
func zd_dkInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "zd_dkInsert",
	})
}
/*对象更新*/
func zd_dkUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "zd_dkUpdate",
	})
}
/*对象删除*/
func zd_dkDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "zd_dkDel",
	})}