package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", RegisterattachmentRoutes)
}

func RegisterattachmentRoutes(route *gin.Engine) {
	 attachmentController := route.Group("attachmentController")
	{
		attachmentController.POST("/insert",attachmentInsert)
		attachmentController.POST("/update",attachmentUpdate)
		attachmentController.GET("/del",attachmentDel)
		attachmentController.GET("/find",attachmentFind)
	}
}

/*对象查询*/
func attachmentFind(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var attachment models.Attachment
		attachments := make([]models.Attachment, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&attachments)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(attachment)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",attachments, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func attachmentInsert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var attachment models.Attachment
	if c.BindJSON(&attachment) == nil {
		affected, err := engine.Insert(attachment)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", attachment,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", attachment,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func attachmentUpdate(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var attachment models.Attachment
	if c.BindJSON(&attachment) == nil {
		affected, err := engine.Id(attachment.Attachment_id).Update(attachment)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", attachment, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", attachment, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func attachmentDel(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var attachment models.Attachment
	if c.Bind(&attachment) == nil {
		affected, err := engine.Id(attachment.Attachment_id).Delete(attachment)
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