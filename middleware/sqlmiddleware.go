package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"../utils"

	"github.com/gin-gonic/gin"
)

//处理jwt
func Sqlmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.ConfigBoolData("app.sqlvalid") {
			c.Next()
			return
		}

		str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)` //此处改为“
		re, err := regexp.Compile(str)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, utils.ResultList{-1, err.Error(), nil, 0, ""})
			return
		}
		//get方式问题
		for k, v := range c.Request.URL.Query() {
			fmt.Println(k)
			validstr := strings.Join(v, ",")
			fmt.Println(re.MatchString(validstr))
			if re.MatchString(validstr) {
				c.AbortWithStatusJSON(http.StatusOK, utils.ResultList{-1, "对不起,请求的内容里面包含非法字符", nil, 0, ""})
				return
			}
		}

		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)
		fmt.Println(re.MatchString(string(x)))
		if re.MatchString(string(x)) {
			c.AbortWithStatusJSON(http.StatusOK, utils.ResultList{-1, "对不起,请求的内容里面包含非法字符", nil, 0, ""})
			return
		}
		c.Next()
	}
}
