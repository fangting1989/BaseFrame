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
	utils.RouterBus.Subscribe("router:register", Registerzd_jzdzbRoutes)
}

func Registerzd_jzdzbRoutes(route *gin.Engine) {
	 zd_jzdzbController := route.Group("zd_jzdzbController")
	{
		zd_jzdzbController.POST("/insert",zd_jzdzbInsert)
		zd_jzdzbController.POST("/update",zd_jzdzbUpdate)
		zd_jzdzbController.GET("/del",zd_jzdzbDel)
		zd_jzdzbController.GET("/find",zd_jzdzbFind)
	}
}

/*对象查询*/
func zd_jzdzbFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var zd_jzdzb models.Zd_jzdzb
		zd_jzdzbs := make([]models.Zd_jzdzb, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&zd_jzdzbs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(zd_jzdzb)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",zd_jzdzbs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func zd_jzdzbInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var zd_jzdzb models.Zd_jzdzb
	if c.BindJSON(&zd_jzdzb) == nil {
		affected, err := engine.Insert(zd_jzdzb)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", zd_jzdzb,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", zd_jzdzb,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func zd_jzdzbUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var zd_jzdzb models.Zd_jzdzb
	if c.BindJSON(&zd_jzdzb) == nil {
		affected, err := engine.Id(zd_jzdzb.ZD_JZDZB_ID).Update(zd_jzdzb)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", zd_jzdzb, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", zd_jzdzb, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func zd_jzdzbDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var zd_jzdzb models.Zd_jzdzb
	if c.Bind(&zd_jzdzb) == nil {
		affected, err := engine.Id(zd_jzdzb.ZD_JZDZB_ID).Delete(zd_jzdzb)
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