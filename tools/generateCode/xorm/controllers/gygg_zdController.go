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
	utils.RouterBus.Subscribe("router:register", Registergygg_zdRoutes)
}

func Registergygg_zdRoutes(route *gin.Engine) {
	 gygg_zdController := route.Group("gygg_zdController")
	{
		gygg_zdController.POST("/insert",gygg_zdInsert)
		gygg_zdController.POST("/update",gygg_zdUpdate)
		gygg_zdController.GET("/del",gygg_zdDel)
		gygg_zdController.GET("/find",gygg_zdFind)
	}
}

/*对象查询*/
func gygg_zdFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var gygg_zd models.Gygg_zd
		gygg_zds := make([]models.Gygg_zd, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&gygg_zds)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(gygg_zd)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",gygg_zds, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func gygg_zdInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var gygg_zd models.Gygg_zd
	if c.BindJSON(&gygg_zd) == nil {
		affected, err := engine.Insert(gygg_zd)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gygg_zd,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gygg_zd,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func gygg_zdUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var gygg_zd models.Gygg_zd
	if c.BindJSON(&gygg_zd) == nil {
		affected, err := engine.Id(gygg_zd.GYGG_ZD_GUID).Update(gygg_zd)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gygg_zd, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gygg_zd, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func gygg_zdDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var gygg_zd models.Gygg_zd
	if c.Bind(&gygg_zd) == nil {
		affected, err := engine.Id(gygg_zd.GYGG_ZD_GUID).Delete(gygg_zd)
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