package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterprojectRoutes)
}

func RegisterprojectRoutes(route *gin.Engine) {
	 projectController := route.Group("projectController")
	{
		projectController.POST("/insert",projectInsert)
		projectController.POST("/update",projectUpdate)
		projectController.GET("/del",projectDel)
		projectController.GET("/find",projectFind)
	}
}

/*对象查询*/
func projectFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var project models.Project
		projects := make([]models.Project, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&projects)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(project)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",projects, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func projectInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var project models.Project
	if c.BindJSON(&project) == nil {
		affected, err := engine.Insert(project)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", project,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", project,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func projectUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var project models.Project
	if c.BindJSON(&project) == nil {
		affected, err := engine.Id(project.PROJECTID).Update(project)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", project, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", project, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func projectDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var project models.Project
	if c.Bind(&project) == nil {
		affected, err := engine.Id(project.PROJECTID).Delete(project)
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