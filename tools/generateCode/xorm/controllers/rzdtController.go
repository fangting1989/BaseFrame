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
	utils.RouterBus.Subscribe("router:register", RegisterrzdtRoutes)
}

func RegisterrzdtRoutes(route *gin.Engine) {
	 rzdtController := route.Group("rzdtController")
	{
		rzdtController.POST("/insert",rzdtInsert)
		rzdtController.POST("/update",rzdtUpdate)
		rzdtController.GET("/del",rzdtDel)
		rzdtController.GET("/find",rzdtFind)
	}
}

/*对象查询*/
func rzdtFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var rzdt models.Rzdt
		rzdts := make([]models.Rzdt, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&rzdts)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(rzdt)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",rzdts, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func rzdtInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var rzdt models.Rzdt
	if c.BindJSON(&rzdt) == nil {
		affected, err := engine.Insert(rzdt)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", rzdt,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", rzdt,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func rzdtUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var rzdt models.Rzdt
	if c.BindJSON(&rzdt) == nil {
		affected, err := engine.Id(rzdt.RZDTID).Update(rzdt)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", rzdt, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", rzdt, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func rzdtDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var rzdt models.Rzdt
	if c.Bind(&rzdt) == nil {
		affected, err := engine.Id(rzdt.RZDTID).Delete(rzdt)
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