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
	"github.com/joshbetz/config"
)

//FieldTypes FieldTypes
var FieldTypes map[string]string

func main() {
	mlog.Start(mlog.LevelInfo, "app.log")
	c := config.New("config.json")
	var value string
	c.Get("name", &value)
	var a = 123
	fmt.Println(a)

	generateStructs()
	generateControllers()

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
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.2.66:3306)/landanalysis_temp?charset=utf8&parseTime=true")

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
	//getCurrentPath
	ret, err := PathExists(getCurrentPath() + "models/")
	check(err)
	if !ret {
		err := os.Mkdir(getCurrentPath()+"models/", os.ModePerm)
		check(err)
	}

	models := tables
	for _, modelName := range models {
		//创建每个表的structs
		res, err := db.Prepare("desc " + modelName)

		check(err)

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

		f, err := os.Create("models/" + strings.Replace(strings.Replace(modelName, "t_", "", 1), "T_", "", 1) + ".go")
		check(err)
		modelName = strings.Title(modelName)

		f.WriteString(`package models
import (
		"time"
		)
type ` + strings.Replace(strings.Replace(modelName, "t_", "", 1), "T_", "", 1) + ` struct {
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
				sql += "not null"
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

	println("code generate finish!!")
	defer db.Close()

}
func check(e error) {
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

func generateControllers() {
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

	Init2()
}

func Init2() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.2.66:3306)/landanalysis_temp?charset=utf8&parseTime=true")

	//查询所有的表
	tbs, err := db.Prepare("show tables ")
	check(err)
	var Tables_in_fsbase string
	var tables2 []string
	qrys, err := tbs.Query()
	for qrys.Next() {
		qrys.Scan(&Tables_in_fsbase)
		glog.V(2).Infoln(Tables_in_fsbase)
		log.Println(Tables_in_fsbase)
		newtables2 := append(tables2, Tables_in_fsbase)
		tables2 = newtables2
	}
	qrys.Close()
	tbs.Close()

	//删除路径
	os.RemoveAll(getCurrentPath() + "controllers/")
	//getCurrentPath
	ret, err := PathExists(getCurrentPath() + "controllers/")
	check(err)
	if !ret {
		err := os.Mkdir(getCurrentPath()+"controllers/", os.ModePerm)
		check(err)
	}
	controllers := tables2
	for _, controllerName := range controllers {
		//创建每个表的controllers
		res, err := db.Prepare("desc " + controllerName)

		check(err)

		qry, err := res.Query()
		check(err)
		//创建文件名
		f, err := os.Create("controllers/" + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + "Controller.go")
		check(err)
		controllerName = strings.Title(controllerName)
		//创建文件内容
		f.WriteString(`package controllers
import (
		"github.com/gin-gonic/gin"
	    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func ` + `Register` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Routes(route *gin.Engine) {
	 ` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Controller := route.Group("` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Controller")
	{
		` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Controller.GET("/insert",` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Insert)
		` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Controller.GET("/update",` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Update)
		` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Controller.GET("/del",` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Del)
		` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Controller.GET("/find",` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Find)
	}
}
/*对象查询*/
func ` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Find(c *gin.Context) {
	//find
	c.JSON(200, gin.H{
		"message": "` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Find",
	})
}
/*对象插入*/
func ` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Insert(c *gin.Context) {
	//insert
	c.JSON(200, gin.H{
		"message": "` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Insert",
	})
}
/*对象更新*/
func ` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Update(c *gin.Context) {
	//update
	c.JSON(200, gin.H{
		"message": "` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Update",
	})
}
/*对象删除*/
func ` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Del(c *gin.Context) {
	//del
	c.JSON(200, gin.H{
		"message": "` + strings.Replace(strings.Replace(controllerName, "t_", "", 1), "T_", "", 1) + `Del",
	})`)
		f.WriteString("}")
		res.Close()
		qry.Close()
		f.Close()
	}

	println("code generate finish!!")
	defer db.Close()
}
