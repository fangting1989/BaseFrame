package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterqybgRoutes)
}

func RegisterqybgRoutes(route *gin.Engine) {
	 qybgController := route.Group("qybgController")
	{
		qybgController.POST("/insert",qybgInsert)
		qybgController.POST("/update",qybgUpdate)
		qybgController.GET("/del",qybgDel)
		qybgController.GET("/find",qybgFind)
	}
}

/*对象查询*/
func qybgFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var qybg models.Qybg
		qybgs := make([]models.Qybg, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&qybgs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(qybg)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",qybgs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func qybgInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var qybg models.Qybg
	if c.BindJSON(&qybg) == nil {
		affected, err := engine.Insert(qybg)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qybg,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qybg,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func qybgUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var qybg models.Qybg
	if c.BindJSON(&qybg) == nil {
		affected, err := engine.Id(qybg.QYBGID).Update(qybg)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qybg, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qybg, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func qybgDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var qybg models.Qybg
	if c.Bind(&qybg) == nil {
		affected, err := engine.Id(qybg.QYBGID).Delete(qybg)
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