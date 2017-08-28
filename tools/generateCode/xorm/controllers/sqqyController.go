package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistersqqyRoutes)
}

func RegistersqqyRoutes(route *gin.Engine) {
	 sqqyController := route.Group("sqqyController")
	{
		sqqyController.POST("/insert",sqqyInsert)
		sqqyController.POST("/update",sqqyUpdate)
		sqqyController.GET("/del",sqqyDel)
		sqqyController.GET("/find",sqqyFind)
	}
}

/*对象查询*/
func sqqyFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var sqqy models.Sqqy
		sqqys := make([]models.Sqqy, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&sqqys)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(sqqy)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",sqqys, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func sqqyInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var sqqy models.Sqqy
	if c.BindJSON(&sqqy) == nil {
		affected, err := engine.Insert(sqqy)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", sqqy,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", sqqy,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func sqqyUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var sqqy models.Sqqy
	if c.BindJSON(&sqqy) == nil {
		affected, err := engine.Id(sqqy.SQQY_ID).Update(sqqy)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", sqqy, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", sqqy, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func sqqyDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var sqqy models.Sqqy
	if c.Bind(&sqqy) == nil {
		affected, err := engine.Id(sqqy.SQQY_ID).Delete(sqqy)
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