package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/olebedev/config"
)

func GetDbConn(dbname string, c *gin.Context) *xorm.Engine {
	cfg, err := config.ParseJsonFile("./config/appconfig.json")
	CheckError(err)
	var connectstring = ""
	switch dbname {
	case "mysql":
		username, err := cfg.String("mysql.username")
		CheckError(err)
		userpwd, err := cfg.String("mysql.userpwd")
		CheckError(err)
		host, err := cfg.String("mysql.host")
		CheckError(err)
		database, err := cfg.String("mysql.database")
		CheckError(err)
		connectstring = username + ":" + userpwd + "@(" + host + ")/" + database + "?charset=utf8"
		break
	}
	engine, err := xorm.NewEngine(dbname, connectstring)
	CheckWebError(err, c)
	engine.SetTableMapper(core.NewPrefixMapper(core.SameMapper{}, "T_"))
	engine.SetColumnMapper(core.SameMapper{})
	return engine
}
