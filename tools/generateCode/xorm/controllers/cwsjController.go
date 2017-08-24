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
	utils.RouterBus.Subscribe("router:register", RegistercwsjRoutes)
}

func RegistercwsjRoutes(route *gin.Engine) {
	 cwsjController := route.Group("cwsjController")
	{
		cwsjController.POST("/insert",cwsjInsert)
		cwsjController.POST("/update",cwsjUpdate)
		cwsjController.GET("/del",cwsjDel)
		cwsjController.GET("/find",cwsjFind)
	}
}

/*对象查询*/
func cwsjFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var cwsj models.Cwsj
		cwsjs := make([]models.Cwsj, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&cwsjs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(cwsj)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",cwsjs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func cwsjInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var cwsj models.Cwsj
	if c.BindJSON(&cwsj) == nil {
		affected, err := engine.Insert(cwsj)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cwsj,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cwsj,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func cwsjUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var cwsj models.Cwsj
	if c.BindJSON(&cwsj) == nil {
		affected, err := engine.Id(cwsj.CWSJ_ID).Update(cwsj)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cwsj, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cwsj, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func cwsjDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var cwsj models.Cwsj
	if c.Bind(&cwsj) == nil {
		affected, err := engine.Id(cwsj.CWSJ_ID).Delete(cwsj)
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