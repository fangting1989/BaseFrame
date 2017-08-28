package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistergfzbRoutes)
}

func RegistergfzbRoutes(route *gin.Engine) {
	 gfzbController := route.Group("gfzbController")
	{
		gfzbController.POST("/insert",gfzbInsert)
		gfzbController.POST("/update",gfzbUpdate)
		gfzbController.GET("/del",gfzbDel)
		gfzbController.GET("/find",gfzbFind)
	}
}

/*对象查询*/
func gfzbFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var gfzb models.Gfzb
		gfzbs := make([]models.Gfzb, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&gfzbs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(gfzb)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",gfzbs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func gfzbInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var gfzb models.Gfzb
	if c.BindJSON(&gfzb) == nil {
		affected, err := engine.Insert(gfzb)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gfzb,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gfzb,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func gfzbUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var gfzb models.Gfzb
	if c.BindJSON(&gfzb) == nil {
		affected, err := engine.Id(gfzb.GFZBID).Update(gfzb)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gfzb, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gfzb, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func gfzbDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var gfzb models.Gfzb
	if c.Bind(&gfzb) == nil {
		affected, err := engine.Id(gfzb.GFZBID).Delete(gfzb)
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