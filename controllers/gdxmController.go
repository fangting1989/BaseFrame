package controllers

import (
	"net/http"

	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegistergdxmRoutes(route *gin.Engine) {
	gdxmController := route.Group("gdxmController")
	{
		gdxmController.POST("/insert", gdxmInsert)
		gdxmController.GET("/update", gdxmUpdate)
		gdxmController.GET("/del", gdxmDel)
		gdxmController.GET("/find", gdxmFind)
		gdxmController.GET("/test", gdxmTest)
	}
}

/*对象查询*/
func gdxmFind(c *gin.Context) {
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		c.JSON(http.StatusOK, gin.H{"errid": 1, errmsg: ""})
		// if pageobject.pagenum == nil {
		// 	pageobject.pagenum = 1
		// }
		// if pageobject.pagesize == nil {
		// 	pageobject.pagesize = 10
		// }
		// var gdxm models.Gdxm
		// gdxms := make([]gdxm, 0)
		// /*书写对应逻辑*/

		// err := engine.Where("1 = 1").Limit(pageobject.pagesize, pageobject.pagesize*(pageobject.pagenum-1)).Find(&gdxms)
		// if err == nil {
		// 	c.JSON(http.StatusOK, gin.H{"error": "not find data"})
		// 	return
		// }
		// total, err := engine.Where("1 = 1").Count(gdxm)
		// if err == nil {
		// 	c.JSON(http.StatusOK, gin.H{"error": "not find data"})
		// 	return
		// }
		// c.JSON(http.StatusOK, gin.H{"errid": 1, errmsg: "", data: gdxms, recordcount: total})
	}

}

/*对象插入*/
func gdxmInsert(c *gin.Context) {
	//insert
	var gdxm models.Gdxm
	// if c.Bind(&gdxm) == nil { BindJSON
	if c.BindJSON(&gdxm) == nil {
		affected, err := engine.Insert(gdxm)
		c.JSON(http.StatusOK, gdxm)
	} else {
		c.JSON(http.StatusOK, gin.H{"errid": -1, errmsg: "no data", data: nil})
	}
}

/*对象更新*/
func gdxmUpdate(c *gin.Context) {
	var gdxm models.Gdxm
	if c.BindJSON(&gdxm) == nil {
		affected, err := engine.Id(gdxm.GDXM_ID).Update(gdxm)
		c.JSON(http.StatusOK, gin.H{"errid": 1, errmsg: "更新成功", data: gdxm})
	} else {
		c.JSON(http.StatusOK, gin.H{"errid": -1, errmsg: "no data", data: nil})
	}
}

/*对象删除*/
func gdxmDel(c *gin.Context) {
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		affected, err := engine.Id(gdxm.GDXM_ID).Delete(gdxm)
		c.JSON(http.StatusOK, gin.H{"errid": 1, errmsg: "删除成功", data: nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"errid": -1, errmsg: "no data", data: nil})
	}
}

func gdxmTest(c *gin.Context) {
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		utils.StructInfo(gdxm)
		c.JSON(http.StatusOK, gin.H{"errid": -1, errmsg: "has data", data: nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"errid": -1, errmsg: "no data", data: nil})
	}

}
