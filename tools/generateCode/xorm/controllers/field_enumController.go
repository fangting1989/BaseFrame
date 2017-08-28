package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", Registerfield_enumRoutes)
}

func Registerfield_enumRoutes(route *gin.Engine) {
	 field_enumController := route.Group("field_enumController")
	{
		field_enumController.POST("/insert",field_enumInsert)
		field_enumController.POST("/update",field_enumUpdate)
		field_enumController.GET("/del",field_enumDel)
		field_enumController.GET("/find",field_enumFind)
	}
}

/*对象查询*/
func field_enumFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var field_enum models.Field_enum
		field_enums := make([]models.Field_enum, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&field_enums)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(field_enum)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",field_enums, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func field_enumInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var field_enum models.Field_enum
	if c.BindJSON(&field_enum) == nil {
		affected, err := engine.Insert(field_enum)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", field_enum,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", field_enum,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func field_enumUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var field_enum models.Field_enum
	if c.BindJSON(&field_enum) == nil {
		affected, err := engine.Id(field_enum.FLD_ENUM_ID).Update(field_enum)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", field_enum, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", field_enum, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func field_enumDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var field_enum models.Field_enum
	if c.Bind(&field_enum) == nil {
		affected, err := engine.Id(field_enum.FLD_ENUM_ID).Delete(field_enum)
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