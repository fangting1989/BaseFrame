package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//FieldTypes FieldTypes
var FieldTypes map[string]string

func main() {
	FieldTypes = map[string]string{
		"bigint":    "int64",
		"int":       "int",
		"tinyint":   "int",
		"smallint":  "int",
		"char":      "string",
		"varchar":   "string",
		"blob":      "[]byte",
		"date":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
		"decimal":   "float64",
		"bit":       "uint64",
	}
	Init()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Init init
func Init() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.120.21:3306)/zqgame_store?charset=utf8&parseTime=true")

	models := [...]string{"appkeys", "apporder", "apps",
		"appstarlevel", "apptype", "apptypejoin", "area", "contentprovider", "datareport",
		"evaluate", "feedback", "indexbannar", "installaction", "ip_area_dictionary",
		"loginaction", "opendetailaction", "regaction", "runaction"}
	check(err)

	for _, modelName := range models {
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
		f, err := os.Create("models/" + modelName + ".go")
		check(err)
		modelName = strings.Title(modelName)

		f.WriteString(`package models
import (
  "time"
)
type ` + modelName + ` struct {
`)

		for qry.Next() {
			qry.Scan(&Field, &Type, &Null, &Key, &DefaultValue, &ExtraValue)
			Title := strings.Title(Field)
			i := strings.Index(Type, "(")
			if i != -1 {
				Type = Type[0:i]
			}

			//Field=Id, Type=smallint, Null=NO, Key=PRI, DefaultValue=, ExtraValue=
			fmt.Printf("Field=%s, Type=%s, Null=%s, Key=%s, DefaultValue=%s, ExtraValue=%s \n", Field, Type, Null, Key, DefaultValue, ExtraValue)
			name := Title
			tp := FieldTypes[Type]
			sql := "`gorm:\"column:" + Field + ";"
			if Null == "NO" {
				sql += "NOT NULL;"
			}
			if Key == "PRI" {
				sql += "PRIMARY KEY;"
			}

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

	println("All Done!! Have Fun!!")
	defer db.Close()

}
