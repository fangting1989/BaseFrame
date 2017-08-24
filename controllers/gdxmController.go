package controllers

import (
	"net/http"

	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistergdxmRoutes)
}

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
	engine := utils.GetDbConn("mysql", c)
	engine.DBMetas()
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
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(gdxm)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "", gdxms, total, utils.ContextToken(c)})
	}
}

/*对象插入*/
func gdxmInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql", c)
	//insert
	var gdxm models.Gdxm
	if c.BindJSON(&gdxm) == nil {
		affected, err := engine.Insert(gdxm)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gdxm, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gdxm, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象更新*/
func gdxmUpdate(c *gin.Context) {
	engine := utils.GetDbConn("mysql", c)
	var gdxm models.Gdxm
	if c.BindJSON(&gdxm) == nil {
		affected, err := engine.Id(gdxm.GD_GUID).Update(gdxm)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gdxm, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gdxm, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func gdxmDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql", c)
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		affected, err := engine.Id(gdxm.GD_GUID).Delete(gdxm)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", nil, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", nil, utils.ContextToken(c)})
		}

	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

func gdxmTest(c *gin.Context) {
	var gdxm models.Gdxm
	if c.Bind(&gdxm) == nil {
		utils.StructInfo(&gdxm)
		c.JSON(http.StatusOK, &utils.ResultList{10, "123", gdxm, 10, ""})
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, ""})
	}

}
