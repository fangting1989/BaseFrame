package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func Registercrgy_kzRoutes(route *gin.Engine) {
	 crgy_kzController := route.Group("crgy_kzController")
	{
		crgy_kzController.GET("/insert",crgy_kzInsert)
		crgy_kzController.GET("/update",crgy_kzUpdate)
		crgy_kzController.GET("/del",crgy_kzDel)
		crgy_kzController.GET("/find",crgy_kzFind)
	}
}
/*对象查询*/
func crgy_kzFind(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "crgy_kzFind",
	})
}
/*对象插入*/
func crgy_kzInsert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "crgy_kzInsert",
	})
}
/*对象更新*/
func crgy_kzUpdate(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "crgy_kzUpdate",
	})
}
/*对象删除*/
func crgy_kzDel(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "crgy_kzDel",
	})}