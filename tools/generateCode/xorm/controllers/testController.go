package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegistertestRoutes)
}

func RegistertestRoutes(route *gin.Engine) {
	 testController := route.Group("testController")
	{
		testController.POST("/insert",testInsert)
		testController.POST("/update",testUpdate)
		testController.GET("/del",testDel)
		testController.GET("/find",testFind)
	}
}

/*对象查询*/
func testFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var test models.Test
		tests := make([]models.Test, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&tests)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(test)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",tests, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func testInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var test models.Test
	if c.BindJSON(&test) == nil {
		affected, err := engine.Insert(test)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", test,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", test,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func testUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var test models.Test
	if c.BindJSON(&test) == nil {
		affected, err := engine.Id(test.ID).Update(test)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", test, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", test, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func testDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var test models.Test
	if c.Bind(&test) == nil {
		affected, err := engine.Id(test.ID).Delete(test)
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