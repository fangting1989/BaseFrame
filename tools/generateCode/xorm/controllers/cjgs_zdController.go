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
	utils.RouterBus.Subscribe("router:register", Registercjgs_zdRoutes)
}

func Registercjgs_zdRoutes(route *gin.Engine) {
	 cjgs_zdController := route.Group("cjgs_zdController")
	{
		cjgs_zdController.POST("/insert",cjgs_zdInsert)
		cjgs_zdController.POST("/update",cjgs_zdUpdate)
		cjgs_zdController.GET("/del",cjgs_zdDel)
		cjgs_zdController.GET("/find",cjgs_zdFind)
	}
}

/*对象查询*/
func cjgs_zdFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var cjgs_zd models.Cjgs_zd
		cjgs_zds := make([]models.Cjgs_zd, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&cjgs_zds)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(cjgs_zd)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",cjgs_zds, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func cjgs_zdInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var cjgs_zd models.Cjgs_zd
	if c.BindJSON(&cjgs_zd) == nil {
		affected, err := engine.Insert(cjgs_zd)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cjgs_zd,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cjgs_zd,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func cjgs_zdUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var cjgs_zd models.Cjgs_zd
	if c.BindJSON(&cjgs_zd) == nil {
		affected, err := engine.Id(cjgs_zd.CJGS_ZD_GUID).Update(cjgs_zd)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", cjgs_zd, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", cjgs_zd, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func cjgs_zdDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var cjgs_zd models.Cjgs_zd
	if c.Bind(&cjgs_zd) == nil {
		affected, err := engine.Id(cjgs_zd.CJGS_ZD_GUID).Delete(cjgs_zd)
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