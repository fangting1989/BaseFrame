package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterenterpriseRoutes)
}

func RegisterenterpriseRoutes(route *gin.Engine) {
	 enterpriseController := route.Group("enterpriseController")
	{
		enterpriseController.POST("/insert",enterpriseInsert)
		enterpriseController.POST("/update",enterpriseUpdate)
		enterpriseController.GET("/del",enterpriseDel)
		enterpriseController.GET("/find",enterpriseFind)
	}
}

/*对象查询*/
func enterpriseFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var enterprise models.Enterprise
		enterprises := make([]models.Enterprise, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&enterprises)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(enterprise)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",enterprises, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func enterpriseInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var enterprise models.Enterprise
	if c.BindJSON(&enterprise) == nil {
		affected, err := engine.Insert(enterprise)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", enterprise,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", enterprise,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func enterpriseUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var enterprise models.Enterprise
	if c.BindJSON(&enterprise) == nil {
		affected, err := engine.Id(enterprise.Enterprise_ID).Update(enterprise)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", enterprise, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", enterprise, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func enterpriseDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var enterprise models.Enterprise
	if c.Bind(&enterprise) == nil {
		affected, err := engine.Id(enterprise.Enterprise_ID).Delete(enterprise)
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