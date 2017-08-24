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
	utils.RouterBus.Subscribe("router:register", RegisterzddxzqRoutes)
}

func RegisterzddxzqRoutes(route *gin.Engine) {
	 zddxzqController := route.Group("zddxzqController")
	{
		zddxzqController.POST("/insert",zddxzqInsert)
		zddxzqController.POST("/update",zddxzqUpdate)
		zddxzqController.GET("/del",zddxzqDel)
		zddxzqController.GET("/find",zddxzqFind)
	}
}

/*对象查询*/
func zddxzqFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var zddxzq models.Zddxzq
		zddxzqs := make([]models.Zddxzq, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&zddxzqs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(zddxzq)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",zddxzqs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func zddxzqInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var zddxzq models.Zddxzq
	if c.BindJSON(&zddxzq) == nil {
		affected, err := engine.Insert(zddxzq)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", zddxzq,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", zddxzq,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func zddxzqUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var zddxzq models.Zddxzq
	if c.BindJSON(&zddxzq) == nil {
		affected, err := engine.Id(zddxzq.ZDDXZQ_ID).Update(zddxzq)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", zddxzq, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", zddxzq, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func zddxzqDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var zddxzq models.Zddxzq
	if c.Bind(&zddxzq) == nil {
		affected, err := engine.Id(zddxzq.ZDDXZQ_ID).Delete(zddxzq)
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