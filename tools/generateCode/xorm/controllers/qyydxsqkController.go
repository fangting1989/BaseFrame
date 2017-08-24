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
	utils.RouterBus.Subscribe("router:register", RegisterqyydxsqkRoutes)
}

func RegisterqyydxsqkRoutes(route *gin.Engine) {
	 qyydxsqkController := route.Group("qyydxsqkController")
	{
		qyydxsqkController.POST("/insert",qyydxsqkInsert)
		qyydxsqkController.POST("/update",qyydxsqkUpdate)
		qyydxsqkController.GET("/del",qyydxsqkDel)
		qyydxsqkController.GET("/find",qyydxsqkFind)
	}
}

/*对象查询*/
func qyydxsqkFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var qyydxsqk models.Qyydxsqk
		qyydxsqks := make([]models.Qyydxsqk, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&qyydxsqks)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(qyydxsqk)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",qyydxsqks, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func qyydxsqkInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var qyydxsqk models.Qyydxsqk
	if c.BindJSON(&qyydxsqk) == nil {
		affected, err := engine.Insert(qyydxsqk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qyydxsqk,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qyydxsqk,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func qyydxsqkUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var qyydxsqk models.Qyydxsqk
	if c.BindJSON(&qyydxsqk) == nil {
		affected, err := engine.Id(qyydxsqk.QYYDXSQKID).Update(qyydxsqk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qyydxsqk, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qyydxsqk, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func qyydxsqkDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var qyydxsqk models.Qyydxsqk
	if c.Bind(&qyydxsqk) == nil {
		affected, err := engine.Id(qyydxsqk.QYYDXSQKID).Delete(qyydxsqk)
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