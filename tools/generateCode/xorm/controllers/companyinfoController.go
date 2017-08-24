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
	utils.RouterBus.Subscribe("router:register", RegistercompanyinfoRoutes)
}

func RegistercompanyinfoRoutes(route *gin.Engine) {
	 companyinfoController := route.Group("companyinfoController")
	{
		companyinfoController.POST("/insert",companyinfoInsert)
		companyinfoController.POST("/update",companyinfoUpdate)
		companyinfoController.GET("/del",companyinfoDel)
		companyinfoController.GET("/find",companyinfoFind)
	}
}

/*对象查询*/
func companyinfoFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var companyinfo models.Companyinfo
		companyinfos := make([]models.Companyinfo, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&companyinfos)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(companyinfo)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",companyinfos, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func companyinfoInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var companyinfo models.Companyinfo
	if c.BindJSON(&companyinfo) == nil {
		affected, err := engine.Insert(companyinfo)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", companyinfo,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", companyinfo,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func companyinfoUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var companyinfo models.Companyinfo
	if c.BindJSON(&companyinfo) == nil {
		affected, err := engine.Id(companyinfo.CompanyInfoID).Update(companyinfo)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", companyinfo, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", companyinfo, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func companyinfoDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var companyinfo models.Companyinfo
	if c.Bind(&companyinfo) == nil {
		affected, err := engine.Id(companyinfo.CompanyInfoID).Delete(companyinfo)
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