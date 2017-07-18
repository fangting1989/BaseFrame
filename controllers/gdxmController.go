package controllers

import (
	"net/http"

	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegistergdxmRoutes(route *gin.Engine) {
	gdxmController := route.Group("gdxmController")
	{
		gdxmController.POST("/insert", gdxmInsert)
		gdxmController.POST("/update", gdxmUpdate)
		gdxmController.GET("/del", gdxmDel)
		gdxmController.GET("/find", gdxmFind)
		gdxmController.GET("/test", gdxmTest)
	}
}

/*对象查询*/
func gdxmFind(c *gin.Context) {
	engine, err := xorm.NewEngine("mysql", "root:123456@192.168.2.66/landanalysis?charset=utf8")
	utils.Check(err)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var gdxm models.Gdxm
		gdxms := make([]models.Gdxm, 0)
		/*书写对应逻辑*/

		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&gdxms)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"error": "not find data"})
			return
		}
		total, err := engine.Where("1 = 1").Count(gdxm)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"error": "not find data"})
			return
		}
		c.JSON(http.StatusOK, &utils.ResultList{1, "", gdxms, total})
	}
}

/*对象插入*/
func gdxmInsert(c *gin.Context) {
	engine, err := xorm.NewEngine("mysql", "root:123456@192.168.2.66/landanalysis?charset=utf8")
	utils.Check(err)
	//insert
	var gdxm models.Gdxm
	if c.BindJSON(&gdxm) == nil {
		affected, err := engine.Insert(gdxm)
		utils.Check(err)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "插入成功", gdxm})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "插入失败", gdxm})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil})
	}
}

/*对象更新*/
func gdxmUpdate(c *gin.Context) {
	engine, err := xorm.NewEngine("mysql", "root:123456@192.168.2.66/landanalysis?charset=utf8")
	utils.Check(err)
	var gdxm models.Gdxm
	if c.BindJSON(&gdxm) == nil {
		affected, err := engine.Id(gdxm.GDXM_ID).Update(gdxm)
		utils.Check(err)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "更新成功", gdxm})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "更新失败", gdxm})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil})
	}
}

/*对象删除*/
func gdxmDel(c *gin.Context) {
	engine, err := xorm.NewEngine("mysql", "root:123456@192.168.2.66/landanalysis?charset=utf8")
	utils.Check(err)
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		affected, err := engine.Id(gdxm.GDXM_ID).Delete(gdxm)
		utils.Check(err)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "删除成功", nil})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "删除失败", nil})
		}

	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil})
	}
}

func gdxmTest(c *gin.Context) {
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		utils.StructInfo(&gdxm)
		c.JSON(http.StatusOK, &utils.ResultList{10, "123", gdxm, 10})
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil})
	}

}
