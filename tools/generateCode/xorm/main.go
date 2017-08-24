package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jbrodriguez/mlog"
	"github.com/olebedev/config"
)

//FieldTypes FieldTypes
var FieldTypes map[string]string

func main() {
	mlog.Start(mlog.LevelInfo, "app.log")
	generateStructs()
}

func generateStructs() {
	FieldTypes = map[string]string{
		"bigint":    "int64",
		"int":       "int",
		"tinyint":   "int",
		"smallint":  "int",
		"char":      "string",
		"varchar":   "string",
		"text":      "string",
		"blob":      "[]uint8",
		"date":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
		"decimal":   "float64",
		"bit":       "uint64",
	}
	Init()
}

//Init init
func Init() {
	db, err := sql.Open("mysql", GetDbConn("mysql"))

	//查询所有的表
	tbs, err := db.Prepare("show tables ")
	check(err)
	var Tables_in_fsbase string
	var tables []string
	qrys, err := tbs.Query()
	for qrys.Next() {
		qrys.Scan(&Tables_in_fsbase)
		glog.V(2).Infoln(Tables_in_fsbase)
		log.Println(Tables_in_fsbase)
		newtables := append(tables, Tables_in_fsbase)
		tables = newtables
	}
	qrys.Close()
	tbs.Close()

	//删除路径
	os.RemoveAll(getCurrentPath() + "models/")
	os.RemoveAll(getCurrentPath() + "controllers/")
	//getCurrentPath
	ret, err := PathExists(getCurrentPath() + "models/")
	ret1, err := PathExists(getCurrentPath() + "controllers/")
	check(err)
	if !ret {
		err := os.Mkdir(getCurrentPath()+"models/", os.ModePerm)
		check(err)
	}
	if !ret1 {
		err := os.Mkdir(getCurrentPath()+"controllers/", os.ModePerm)
		check(err)
	}
	models := tables
	controllers := tables

	for _, modelName := range models {
		//创建每个表的structs
		res, err := db.Prepare("desc " + modelName)
		check(err)
		var newMname = strings.Replace(strings.Replace(modelName, "t_", "", 1), "T_", "", 1)
		qry, err := res.Query()
		check(err)
		// Field    | Type        | Null | Key | Default | Extra
		var (
			Field        string
			Type         string
			Null         string
			Key          string
			DefaultValue string
			ExtraValue   string
		)

		f, err := os.Create("models/" + newMname + ".go")
		check(err)
		modelName = strings.Title(modelName)

		f.WriteString(`package models
import (
		"time"
		)
type ` + strFirstToUpper(newMname) + ` struct {
`)
		for qry.Next() {
			qry.Scan(&Field, &Type, &Null, &Key, &DefaultValue, &ExtraValue)
			SourceType := Type
			Title := strings.Title(Field)
			i := strings.Index(Type, "(")
			if i != -1 {
				Type = Type[0:i]
			}

			//Field=Id, Type=smallint, Null=NO, Key=PRI, DefaultValue=, ExtraValue=
			fmt.Printf("Field=%s, Type=%s, Null=%s, Key=%s, DefaultValue=%s, ExtraValue=%s \n", Field, Type, Null, Key, DefaultValue, ExtraValue)
			name := Title
			tp := FieldTypes[Type]
			sql := "`xorm:\""
			if Null == "NO" {
				sql += " not null "
			}
			if Key == "PRI" {
				sql += " pk "
			}
			sql += strings.ToUpper(SourceType)

			sql += "\"`"
			line := fmt.Sprintf("  %-10s\t%-10s\t%-20s", name, tp, sql)
			f.WriteString(line + "\n")
			Field, Type, Null, Key, DefaultValue, ExtraValue = "", "", "", "", "", ""
		}
		f.WriteString("}")
		res.Close()
		qry.Close()
		f.Close()
	}

	for _, controllerName := range controllers {
		//创建每个表的controllers
		res, err := db.Prepare("desc " + controllerName)
		check(err)
		var newCname = strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1)
		qry, err := res.Query()
		check(err)
		var (
			Field        string
			Type         string
			Null         string
			Key          string
			DefaultValue string
			ExtraValue   string
			Pr           string
		)

		//创建文件名
		f, err := os.Create("controllers/" + newCname + "Controller.go")
		check(err)
		controllerName = strings.Title(controllerName)
		Pr = ""
		for qry.Next() {
			qry.Scan(&Field, &Type, &Null, &Key, &DefaultValue, &ExtraValue)
			if Pr == "" {
				Pr = Field
			}
			if Key == "PRI" {
				Pr = Field
			}
			Field, Type, Null, Key, DefaultValue, ExtraValue = "", "", "", "", "", ""

		}

		//创建文件内容
		f.WriteString(`package controllers
		
import (
	"net/http"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	utils.RouterBus.Subscribe("router:register", Register` + newCname + `Routes)
}

func Register` + newCname + `Routes(route *gin.Engine) {
	 ` + newCname + `Controller := route.Group("` + newCname + `Controller")
	{
		` + newCname + `Controller.POST("/insert",` + newCname + `Insert)
		` + newCname + `Controller.POST("/update",` + newCname + `Update)
		` + newCname + `Controller.GET("/del",` + newCname + `Del)
		` + newCname + `Controller.GET("/find",` + newCname + `Find)
	}
}

/*对象查询*/
func ` + newCname + `Find(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var pageobject utils.Pageobject
	if c.Bind(&pageobject) == nil {
		if pageobject.Pagenum == 0 {
			pageobject.Pagenum = 1
		}
		if pageobject.Pagesize == 0 {
			pageobject.Pagesize = 10
		}
		var ` + newCname + ` models.` + strFirstToUpper(newCname) + `
		` + newCname + `s := make([]models.` + strFirstToUpper(newCname) + `, 0)
		/*书写对应逻辑*/
		err := engine.Where("1 = 1").Limit(pageobject.Pagesize, pageobject.Pagesize*(pageobject.Pagenum-1)).Find(&` + newCname + `s)
		utils.CheckWebError(err, c)
		total, err := engine.Where("1 = 1").Count(` + newCname + `)
		utils.CheckWebError(err, c)
		c.JSON(http.StatusOK, &utils.ResultList{1, "",` + newCname + `s, total,utils.ContextToken(c)})
	}
}

/*对象插入*/
func ` + newCname + `Insert(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	//insert
	var ` + newCname + ` models.` + strFirstToUpper(newCname) +
			`
	if c.BindJSON(&` + newCname + `) == nil {
		affected, err := engine.Insert(` + newCname + `)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", ` + newCname + `,utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", ` + newCname + `,utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil,utils.ContextToken(c)})
	}
}

/*对象更新*/
func ` + newCname + `Update(c *gin.Context) {
	engine:= utils.GetDbConn("mysql",c)
	var ` + newCname + ` models.` + strFirstToUpper(newCname) +
			`
	if c.BindJSON(&` + newCname + `) == nil {
		affected, err := engine.Id(` + newCname + `.` + Pr + `).Update(` + newCname + `)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", ` + newCname + `, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", ` + newCname + `, utils.ContextToken(c)})
		}
	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
}

/*对象删除*/
func ` + newCname + `Del(c *gin.Context) {
	engine := utils.GetDbConn("mysql",c)
	var ` + newCname + ` models.` + strFirstToUpper(newCname) +
			`
	if c.Bind(&` + newCname + `) == nil {
		affected, err := engine.Id(` + newCname + `.` + Pr + `).Delete(` + newCname + `)
		utils.CheckWebError(err, c)
		if affected > 0 {
			c.JSON(http.StatusOK, &utils.ResultObject{1, "操作成功", nil, utils.ContextToken(c)})
		} else {
			c.JSON(http.StatusOK, &utils.ResultObject{-1, "操作失败", nil, utils.ContextToken(c)})
		}

	} else {
		c.JSON(http.StatusOK, &utils.ResultObject{-1, "参数解析错误", nil, utils.ContextToken(c)})
	}
`)
		f.WriteString("}")
		res.Close()
		qry.Close()
		f.Close()
	}

	println("code generate finish!!")
	defer db.Close()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	check(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//字符串首字母小写转换成大写
func strFirstToUpper(str string) string {
	var newStr string
	vv := []rune(str)
	if vv[0] >= 97 && vv[0] <= 122 {
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				newStr += string(vv[i]) // + string(vv[i+1])
			} else {
				newStr += string(vv[i])
			}
		}
		return newStr
	} else {
		return str
	}
}

func GetDbConn(dbname string) string {
	//root:123456@tcp(192.168.2.66:3306)/landanalysis_temp?charset=utf8&parseTime=true
	cfg, err := config.ParseJsonFile("./config.json")
	Check(err)
	var connectstring = ""
	switch dbname {
	case "mysql":
		username, err := cfg.String("mysql.username")
		Check(err)
		userpwd, err := cfg.String("mysql.userpwd")
		Check(err)
		host, err := cfg.String("mysql.host")
		Check(err)
		database, err := cfg.String("mysql.database")
		Check(err)
		port, err := cfg.String("mysql.port")
		Check(err)
		connectstring = username + ":" + userpwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=true"
		break
	}
	return connectstring
}
