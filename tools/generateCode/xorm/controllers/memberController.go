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
	utils.RouterBus.Subscribe("router:register", RegistermemberRoutes)
}

func RegistermemberRoutes(route *gin.Engine) {
	 memberController := route.Group("memberController")
	{
		memberController.POST("/insert",memberInsert)
		memberController.POST("/update",memberUpdate)
		memberController.GET("/del",memberDel)
		memberController.GET("/find",memberFind)
	}
}

/*对象查询*/
func memberFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var member models.Member
		members := make([]models.Member, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&members)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(member)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",members, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func memberInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var member models.Member
	if c.BindJSON(&member) == nil {
		affected, err := engine.Insert(member)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", member,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", member,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func memberUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var member models.Member
	if c.BindJSON(&member) == nil {
		affected, err := engine.Id(member.Member_ID).Update(member)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", member, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", member, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func memberDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var member models.Member
	if c.Bind(&member) == nil {
		affected, err := engine.Id(member.Member_ID).Delete(member)
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