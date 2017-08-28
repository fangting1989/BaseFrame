package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistercjgsRoutes)
}

func RegistercjgsRoutes(route *gin.Engine) {
	 cjgsController := route.Group("cjgsController")
	{
		cjgsController.POST("/insert",cjgsInsert)
		cjgsController.POST("/update",cjgsUpdate)
		cjgsController.GET("/del",cjgsDel)
		cjgsController.GET("/find",cjgsFind)
	}
}

/*对象查询*/
func cjgsFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var cjgs models.Cjgs
		cjgss := make([]models.Cjgs, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&cjgss)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(cjgs)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",cjgss, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func cjgsInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var cjgs models.Cjgs
	if c.BindJSON(&cjgs) == nil {
		affected, err := engine.Insert(cjgs)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cjgs,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cjgs,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func cjgsUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var cjgs models.Cjgs
	if c.BindJSON(&cjgs) == nil {
		affected, err := engine.Id(cjgs.CJGS_GUID).Update(cjgs)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cjgs, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cjgs, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func cjgsDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var cjgs models.Cjgs
	if c.Bind(&cjgs) == nil {
		affected, err := engine.Id(cjgs.CJGS_GUID).Delete(cjgs)
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