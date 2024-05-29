package jwt

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	token, err := GetToken(JwtPayLoad{
		Id:      1,
		Account: "20212021",
	}, "default-secret", 3600)
	t.Log(token, err)
}

func TestParseToken(t *testing.T) {
	parseToken, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiYWNjb3VudCI6IjIwMjEyMDIxIiwiZXhwIjoxNzI5OTQzNjM2fQ.KsS53bvXi--krD0g-1IyPmqfwodnmxMScivLcCqJTCQ",
		"default-secret", 3600)
	fmt.Println(parseToken, err)
}
