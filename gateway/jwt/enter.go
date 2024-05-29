package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var (
	JwtSecret = ""
	JwtExpire = 0
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	Id      uint64 `json:"id"`
	Account string `json:"account"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

func init() {
	JwtSecret = viper.GetString("jwt-secret")
	JwtExpire = viper.GetInt("jwt-expiration")

	fmt.Println("jwt: ", JwtSecret, JwtExpire)
}
