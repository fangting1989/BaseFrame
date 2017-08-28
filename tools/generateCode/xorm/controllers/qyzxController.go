package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterqyzxRoutes)
}

func RegisterqyzxRoutes(route *gin.Engine) {
	 qyzxController := route.Group("qyzxController")
	{
		qyzxController.POST("/insert",qyzxInsert)
		qyzxController.POST("/update",qyzxUpdate)
		qyzxController.GET("/del",qyzxDel)
		qyzxController.GET("/find",qyzxFind)
	}
}

/*对象查询*/
func qyzxFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var qyzx models.Qyzx
		qyzxs := make([]models.Qyzx, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&qyzxs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(qyzx)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",qyzxs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func qyzxInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var qyzx models.Qyzx
	if c.BindJSON(&qyzx) == nil {
		affected, err := engine.Insert(qyzx)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qyzx,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qyzx,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func qyzxUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var qyzx models.Qyzx
	if c.BindJSON(&qyzx) == nil {
		affected, err := engine.Id(qyzx.QYZXID).Update(qyzx)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", qyzx, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", qyzx, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func qyzxDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var qyzx models.Qyzx
	if c.Bind(&qyzx) == nil {
		affected, err := engine.Id(qyzx.QYZXID).Delete(qyzx)
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