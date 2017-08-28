package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", Registercrgy_kzRoutes)
}

func Registercrgy_kzRoutes(route *gin.Engine) {
	 crgy_kzController := route.Group("crgy_kzController")
	{
		crgy_kzController.POST("/insert",crgy_kzInsert)
		crgy_kzController.POST("/update",crgy_kzUpdate)
		crgy_kzController.GET("/del",crgy_kzDel)
		crgy_kzController.GET("/find",crgy_kzFind)
	}
}

/*对象查询*/
func crgy_kzFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var crgy_kz models.Crgy_kz
		crgy_kzs := make([]models.Crgy_kz, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&crgy_kzs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(crgy_kz)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",crgy_kzs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func crgy_kzInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var crgy_kz models.Crgy_kz
	if c.BindJSON(&crgy_kz) == nil {
		affected, err := engine.Insert(crgy_kz)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", crgy_kz,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", crgy_kz,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func crgy_kzUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var crgy_kz models.Crgy_kz
	if c.BindJSON(&crgy_kz) == nil {
		affected, err := engine.Id(crgy_kz.CRKZ_GUID).Update(crgy_kz)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", crgy_kz, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", crgy_kz, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func crgy_kzDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var crgy_kz models.Crgy_kz
	if c.Bind(&crgy_kz) == nil {
		affected, err := engine.Id(crgy_kz.CRKZ_GUID).Delete(crgy_kz)
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