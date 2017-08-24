package middleware

import (
	"net/http"
	"time"

	"../utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//处理jwt
func Jwtmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.JwtMode() {
			c.Next()
			return
		}
		TokenString := c.GetHeader("Token")
		if TokenString == "" {
			c.AbortWithStatusJSON(http.StatusOK, utils.ResultList{-1, "用户标识错误", nil, 0, ""})
			return
		}
		websitejwtkey := []byte("$$%%%$$")
		var mobj = decryptjwt(TokenString, websitejwtkey)
		if mobj.ExpiresAt == 0 {
			c.AbortWithStatusJSON(http.StatusOK, utils.ResultList{-1, "用户标识无法解析", nil, 0, ""})
			return
		}
		//判断token
		// tokenExpires := utils.ConfigIntData("app.tokenExpires")
		// if tokenExpires == 0 {
		// 	tokenExpires = 2
		// }
		// tm := time.Unix(mobj.ExpiresAt, 0)
		// now := time.Now()
		// subM := now.Sub(tm)
		// if (int)(subM.Minutes()) > (tokenExpires + 2) {
		// 	//过期
		// 	c.JSON(http.StatusOK, utils.ResultList{-1, "用户授权过期", nil, 0, ""})
		// 	c.Abort()
		// 	return
		// }
		//重新分配
		ss, err := encryptjwt(mobj.Issuer, mobj.Id, websitejwtkey)
		if err != nil {
			c.JSON(http.StatusOK, utils.ResultList{-1, err.Error(), nil, 0, ""})
			c.Abort()
			return
		}
		//设置tokenstring
		c.Set("TokenString", ss)
		c.Next()
		return
	}
}
func encryptjwt(username string, usercode string, key []byte) (string, error) {
	timestamp := time.Now().Unix()
	claims := &jwt.StandardClaims{
		ExpiresAt: timestamp,
		Issuer:    username,
		Id:        usercode,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
func decryptjwt(token string, key []byte) jwt.StandardClaims {
	mclaims := jwt.StandardClaims{}
	jwt.ParseWithClaims(token, &mclaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return mclaims
}
