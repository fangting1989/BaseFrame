package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	// db, err := gorm.Open("mysql", "root:123456@/gotest?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// var T_member models.T_member

	// db.Last(&T_member)

	// defer db.Close()

	c.JSON(200, gin.H{
		"message": "123",
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
