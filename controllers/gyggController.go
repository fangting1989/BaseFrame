package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistergyggRoutes)
}

func RegistergyggRoutes(route *gin.Engine) {
	 gyggController := route.Group("gyggController")
	{
		gyggController.POST("/insert",gyggInsert)
		gyggController.POST("/update",gyggUpdate)
		gyggController.GET("/del",gyggDel)
		gyggController.GET("/find",gyggFind)
	}
}

/*对象查询*/
func gyggFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var gygg models.Gygg
		gyggs := make([]models.Gygg, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&gyggs)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(gygg)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",gyggs, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func gyggInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var gygg models.Gygg
	if c.BindJSON(&gygg) == nil {
		affected, err := engine.Insert(gygg)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gygg,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gygg,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func gyggUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var gygg models.Gygg
	if c.BindJSON(&gygg) == nil {
		affected, err := engine.Id(gygg.GYGG_GUID).Update(gygg)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", gygg, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", gygg, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func gyggDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var gygg models.Gygg
	if c.Bind(&gygg) == nil {
		affected, err := engine.Id(gygg.GYGG_GUID).Delete(gygg)
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