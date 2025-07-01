package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenValid            = errors.New("未知错误")
	TokenExpired          = errors.New("token已过期")
	TokenNotValidYet      = errors.New("token尚未激活")
	TokenMalformed        = errors.New("这不是一个token")
	TokenSignatureInvalid = errors.New("无效签名")
	TokenInvalid          = errors.New("无法处理此token")
)

type JwtPayLoad struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
}
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

type JWT struct {
	Signed []byte
}
