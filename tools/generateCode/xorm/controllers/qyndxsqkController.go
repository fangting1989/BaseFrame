package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterqyndxsqkRoutes)
}

func RegisterqyndxsqkRoutes(route *gin.Engine) {
	 qyndxsqkController := route.Group("qyndxsqkController")
	{
		qyndxsqkController.POST("/insert",qyndxsqkInsert)
		qyndxsqkController.POST("/update",qyndxsqkUpdate)
		qyndxsqkController.GET("/del",qyndxsqkDel)
		qyndxsqkController.GET("/find",qyndxsqkFind)
	}
}

/*对象查询*/
func qyndxsqkFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var qyndxsqk models.Qyndxsqk
		qyndxsqks := make([]models.Qyndxsqk, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&qyndxsqks)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(qyndxsqk)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",qyndxsqks, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func qyndxsqkInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var qyndxsqk models.Qyndxsqk
	if c.BindJSON(&qyndxsqk) == nil {
		affected, err := engine.Insert(qyndxsqk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qyndxsqk,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qyndxsqk,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func qyndxsqkUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var qyndxsqk models.Qyndxsqk
	if c.BindJSON(&qyndxsqk) == nil {
		affected, err := engine.Id(qyndxsqk.QYNDXSQKID).Update(qyndxsqk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qyndxsqk, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qyndxsqk, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func qyndxsqkDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var qyndxsqk models.Qyndxsqk
	if c.Bind(&qyndxsqk) == nil {
		affected, err := engine.Id(qyndxsqk.QYNDXSQKID).Delete(qyndxsqk)
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