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
	utils.RouterBus.Subscribe("router:register", RegistercompanyRoutes)
}

func RegistercompanyRoutes(route *gin.Engine) {
	 companyController := route.Group("companyController")
	{
		companyController.POST("/insert",companyInsert)
		companyController.POST("/update",companyUpdate)
		companyController.GET("/del",companyDel)
		companyController.GET("/find",companyFind)
	}
}

/*对象查询*/
func companyFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var company models.Company
		companys := make([]models.Company, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&companys)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(company)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",companys, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func companyInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var company models.Company
	if c.BindJSON(&company) == nil {
		affected, err := engine.Insert(company)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", company,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", company,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func companyUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var company models.Company
	if c.BindJSON(&company) == nil {
		affected, err := engine.Id(company.Company_ID).Update(company)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", company, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", company, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func companyDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var company models.Company
	if c.Bind(&company) == nil {
		affected, err := engine.Id(company.Company_ID).Delete(company)
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