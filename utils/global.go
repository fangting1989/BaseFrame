package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olebedev/config"
)

type Pageobject struct {
	Pagenum  int `json:"pagenum"`
	Pagesize int `json:"pagesize"`
}

type ResultList struct {
	Errid       int         `json:"errid"`
	Errmsg      string      `json:"errmsg"`
	Data        interface{} `json:"data"`
	Recordcount int64       `json:"recordcount"`
	Token       string      `json:"token"`
}

type ResultObject struct {
	Errid  int         `json:"errid"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}

func StructInfo(obj interface{}) string {
	val := reflect.ValueOf(obj).Elem()
	// fmt.Println("type:", t.Name())

	// if t.Kind() != reflect.Struct {
	// 	fmt.Println("err:type invalid")
	// }
	// s := t.Elem()
	// typeOfT := s.Type() //把s.Type()返回的Type对象复制给typeofT，typeofT也是一个反射
	for k := 0; k < val.NumField(); k++ {

		valueField := val.Field(k)
		typeField := val.Type().Field(k)
		tag := typeField.Tag

		// var aa = 123
		// f := val.Field(k)                                                                                                               //迭代s的各个域，注意每个域仍然是反射。
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name")) //提取了每个域的名字

		// fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
	}
	return "123"
}

func CheckError(e error) {
	if e != nil {
		LogInfo(e.Error(), "error")
		panic(e)
	}
}

func CheckWebError(e error, c *gin.Context) {
	if e != nil {
		LogInfo(e.Error(), "error")
		c.JSON(http.StatusOK, ResultList{-1, e.Error(), nil, 0, ContextToken(c)})
		panic(e)
	}
}

func StartPort() string {
	var startport = "8000"
	cfg, err := config.ParseJsonFile("./config/appconfig.json")
	if err != nil {
		return startport
	}
	newstartport, err := cfg.String("app.startport")
	if err != nil {
		return startport
	} else {
		return newstartport
	}
}

func JwtMode() bool {
	cfg, err := config.ParseJsonFile("./config/appconfig.json")
	if err != nil {
		return false
	}
	nmode, err := cfg.String("app.jwtmode")
	if err != nil {
		return false
	} else {
		b, err := strconv.ParseBool(nmode)
		if err != nil {
			return false
		} else {
			return b
		}
	}
}

func ConfigIntData(name string) int {
	cfg, err := config.ParseJsonFile("./config/appconfig.json")
	if err != nil {
		return 0
	}
	nmode, err := cfg.String(name)
	if err != nil {
		return 0
	} else {
		b, err := strconv.Atoi(nmode)
		if err != nil {
			return 0
		} else {
			return b
		}
	}
}

func ConfigBoolData(name string) bool {
	cfg, err := config.ParseJsonFile("./config/appconfig.json")
	if err != nil {
		return false
	}
	nmode, err := cfg.String(name)
	if err != nil {
		return false
	} else {
		b, err := strconv.ParseBool(nmode)
		if err != nil {
			return false
		} else {
			return b
		}
	}
}

func ContextKeyValue(keyName string, c *gin.Context) string {
	val, has := c.Get(keyName)
	if !has {
		return ""
	} else {
		return val.(string)
	}
}

func ContextToken(c *gin.Context) string {
	val, has := c.Get("TokenString")
	if !has {
		return ""
	} else {
		return val.(string)
	}
}
