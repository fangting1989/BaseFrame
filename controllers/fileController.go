package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterfileRoutes)
}

func RegisterfileRoutes(route *gin.Engine) {
	 fileController := route.Group("fileController")
	{
		fileController.POST("/insert",fileInsert)
		fileController.POST("/update",fileUpdate)
		fileController.GET("/del",fileDel)
		fileController.GET("/find",fileFind)
	}
}

/*对象查询*/
func fileFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var file models.File
		files := make([]models.File, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&files)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(file)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",files, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func fileInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var file models.File
	if c.BindJSON(&file) == nil {
		affected, err := engine.Insert(file)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", file,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", file,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func fileUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var file models.File
	if c.BindJSON(&file) == nil {
		affected, err := engine.Id(file.File_id).Update(file)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", file, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", file, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func fileDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var file models.File
	if c.Bind(&file) == nil {
		affected, err := engine.Id(file.File_id).Delete(file)
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