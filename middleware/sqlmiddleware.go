package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

//处理jwt
func Sqlmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if !utils.JwtMode() {
		// 	c.Next()
		// 	return
		// }
		//fmt.Println()
		for k, v := range c.Request.URL.Query() {
			fmt.Printf(k)
			fmt.Printf(strings.Join(v, ","))
		}

	}
}
