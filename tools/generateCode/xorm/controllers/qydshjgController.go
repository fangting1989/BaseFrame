package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterqydshjgRoutes)
}

func RegisterqydshjgRoutes(route *gin.Engine) {
	 qydshjgController := route.Group("qydshjgController")
	{
		qydshjgController.POST("/insert",qydshjgInsert)
		qydshjgController.POST("/update",qydshjgUpdate)
		qydshjgController.GET("/del",qydshjgDel)
		qydshjgController.GET("/find",qydshjgFind)
	}
}

/*对象查询*/
func qydshjgFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var qydshjg models.Qydshjg
		qydshjgs := make([]models.Qydshjg, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&qydshjgs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(qydshjg)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",qydshjgs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func qydshjgInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var qydshjg models.Qydshjg
	if c.BindJSON(&qydshjg) == nil {
		affected, err := engine.Insert(qydshjg)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qydshjg,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qydshjg,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func qydshjgUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var qydshjg models.Qydshjg
	if c.BindJSON(&qydshjg) == nil {
		affected, err := engine.Id(qydshjg.QYDSHJGID).Update(qydshjg)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qydshjg, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qydshjg, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func qydshjgDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var qydshjg models.Qydshjg
	if c.Bind(&qydshjg) == nil {
		affected, err := engine.Id(qydshjg.QYDSHJGID).Delete(qydshjg)
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