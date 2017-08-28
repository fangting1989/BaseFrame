package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistergzqyRoutes)
}

func RegistergzqyRoutes(route *gin.Engine) {
	 gzqyController := route.Group("gzqyController")
	{
		gzqyController.POST("/insert",gzqyInsert)
		gzqyController.POST("/update",gzqyUpdate)
		gzqyController.GET("/del",gzqyDel)
		gzqyController.GET("/find",gzqyFind)
	}
}

/*对象查询*/
func gzqyFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var gzqy models.Gzqy
		gzqys := make([]models.Gzqy, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&gzqys)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(gzqy)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",gzqys, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func gzqyInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var gzqy models.Gzqy
	if c.BindJSON(&gzqy) == nil {
		affected, err := engine.Insert(gzqy)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gzqy,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gzqy,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func gzqyUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var gzqy models.Gzqy
	if c.BindJSON(&gzqy) == nil {
		affected, err := engine.Id(gzqy.GZQY_ID).Update(gzqy)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gzqy, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gzqy, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func gzqyDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var gzqy models.Gzqy
	if c.Bind(&gzqy) == nil {
		affected, err := engine.Id(gzqy.GZQY_ID).Delete(gzqy)
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