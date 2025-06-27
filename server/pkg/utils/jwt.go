package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type JwtPayLoad struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
}
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

func getToken() (t string, err error) {
	// CustomClaims因为继承(不是)jwt.RegisteredClaims的结构因此满足Claims接口
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		JwtPayLoad: JwtPayLoad{},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})
	var Key = []byte("AnyKey")
	// 自己定义密钥 生成不一样的签名
	str, err := token.SignedString(Key)
	if err != nil {
		return
	}
	return str, err
}

func parseToken(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AnyKey"), nil // 此处返回参数interface{}是自己的密钥
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}
