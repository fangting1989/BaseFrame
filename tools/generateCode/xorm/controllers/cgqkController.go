package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistercgqkRoutes)
}

func RegistercgqkRoutes(route *gin.Engine) {
	 cgqkController := route.Group("cgqkController")
	{
		cgqkController.POST("/insert",cgqkInsert)
		cgqkController.POST("/update",cgqkUpdate)
		cgqkController.GET("/del",cgqkDel)
		cgqkController.GET("/find",cgqkFind)
	}
}

/*对象查询*/
func cgqkFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var cgqk models.Cgqk
		cgqks := make([]models.Cgqk, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&cgqks)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(cgqk)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",cgqks, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func cgqkInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var cgqk models.Cgqk
	if c.BindJSON(&cgqk) == nil {
		affected, err := engine.Insert(cgqk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cgqk,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cgqk,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func cgqkUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var cgqk models.Cgqk
	if c.BindJSON(&cgqk) == nil {
		affected, err := engine.Id(cgqk.CGQKID).Update(cgqk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cgqk, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cgqk, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func cgqkDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var cgqk models.Cgqk
	if c.Bind(&cgqk) == nil {
		affected, err := engine.Id(cgqk.CGQKID).Delete(cgqk)
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