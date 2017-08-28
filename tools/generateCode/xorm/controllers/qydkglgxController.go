package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterqydkglgxRoutes)
}

func RegisterqydkglgxRoutes(route *gin.Engine) {
	 qydkglgxController := route.Group("qydkglgxController")
	{
		qydkglgxController.POST("/insert",qydkglgxInsert)
		qydkglgxController.POST("/update",qydkglgxUpdate)
		qydkglgxController.GET("/del",qydkglgxDel)
		qydkglgxController.GET("/find",qydkglgxFind)
	}
}

/*对象查询*/
func qydkglgxFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var qydkglgx models.Qydkglgx
		qydkglgxs := make([]models.Qydkglgx, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&qydkglgxs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(qydkglgx)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",qydkglgxs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func qydkglgxInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var qydkglgx models.Qydkglgx
	if c.BindJSON(&qydkglgx) == nil {
		affected, err := engine.Insert(qydkglgx)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qydkglgx,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qydkglgx,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func qydkglgxUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var qydkglgx models.Qydkglgx
	if c.BindJSON(&qydkglgx) == nil {
		affected, err := engine.Id(qydkglgx.QYDKGLGXID).Update(qydkglgx)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qydkglgx, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qydkglgx, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func qydkglgxDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var qydkglgx models.Qydkglgx
	if c.Bind(&qydkglgx) == nil {
		affected, err := engine.Id(qydkglgx.QYDKGLGXID).Delete(qydkglgx)
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