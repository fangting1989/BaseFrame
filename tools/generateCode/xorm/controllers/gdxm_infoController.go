package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", Registergdxm_infoRoutes)
}

func Registergdxm_infoRoutes(route *gin.Engine) {
	 gdxm_infoController := route.Group("gdxm_infoController")
	{
		gdxm_infoController.POST("/insert",gdxm_infoInsert)
		gdxm_infoController.POST("/update",gdxm_infoUpdate)
		gdxm_infoController.GET("/del",gdxm_infoDel)
		gdxm_infoController.GET("/find",gdxm_infoFind)
	}
}

/*对象查询*/
func gdxm_infoFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var gdxm_info models.Gdxm_info
		gdxm_infos := make([]models.Gdxm_info, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&gdxm_infos)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(gdxm_info)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",gdxm_infos, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func gdxm_infoInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var gdxm_info models.Gdxm_info
	if c.BindJSON(&gdxm_info) == nil {
		affected, err := engine.Insert(gdxm_info)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gdxm_info,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gdxm_info,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func gdxm_infoUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var gdxm_info models.Gdxm_info
	if c.BindJSON(&gdxm_info) == nil {
		affected, err := engine.Id(gdxm_info.GDXM_INFOID).Update(gdxm_info)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gdxm_info, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gdxm_info, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func gdxm_infoDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var gdxm_info models.Gdxm_info
	if c.Bind(&gdxm_info) == nil {
		affected, err := engine.Id(gdxm_info.GDXM_INFOID).Delete(gdxm_info)
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