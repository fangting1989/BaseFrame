package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", Registerzd_dkRoutes)
}

func Registerzd_dkRoutes(route *gin.Engine) {
	 zd_dkController := route.Group("zd_dkController")
	{
		zd_dkController.POST("/insert",zd_dkInsert)
		zd_dkController.POST("/update",zd_dkUpdate)
		zd_dkController.GET("/del",zd_dkDel)
		zd_dkController.GET("/find",zd_dkFind)
	}
}

/*对象查询*/
func zd_dkFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var zd_dk models.Zd_dk
		zd_dks := make([]models.Zd_dk, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&zd_dks)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(zd_dk)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",zd_dks, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func zd_dkInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var zd_dk models.Zd_dk
	if c.BindJSON(&zd_dk) == nil {
		affected, err := engine.Insert(zd_dk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", zd_dk,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", zd_dk,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func zd_dkUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var zd_dk models.Zd_dk
	if c.BindJSON(&zd_dk) == nil {
		affected, err := engine.Id(zd_dk.JZDSX_GUID).Update(zd_dk)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", zd_dk, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", zd_dk, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func zd_dkDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var zd_dk models.Zd_dk
	if c.Bind(&zd_dk) == nil {
		affected, err := engine.Id(zd_dk.JZDSX_GUID).Delete(zd_dk)
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