package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/common"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/jwtutils"
)

func main() {
	jwtGen, err := jwtutils.NewJWTGenerator("./test.private.pem")
	common.HandleError(err)

	tokenStr, err := jwtGen.GenerateJWT(jwt.MapClaims{
		"sub":  "1234",
		"name": "test1234",
	})
	common.HandleError(err)

	println(tokenStr)
}
