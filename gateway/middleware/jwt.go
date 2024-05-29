package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-note/gateway/jwt"
	"strings"
)

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization header
		tokenString := c.GetHeader("Authorization")
		splits := strings.Split(tokenString, " ")
		if len(splits) != 2 || splits[0] != "Bearer" {
			c.JSON(401, gin.H{"msg": "please provide your jwt token"})
			c.Abort()
			return
		}
		tokenString = splits[1]
		// parse token
		_, err := jwt.ParseToken(tokenString, jwt.JwtSecret, int64(jwt.JwtExpire))
		if err != nil {
			c.JSON(401, gin.H{"msg": "invalid token, please login again"})
			c.Abort()
			return
		}
		c.Next()
	}
}
