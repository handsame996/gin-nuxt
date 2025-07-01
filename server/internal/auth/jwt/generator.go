package jwt

import (
	"example/template/internal/configs"
	"example/template/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewJWT(signed string) *JWT {
	return &JWT{
		Signed: []byte(signed),
	}
}

func (j *JWT) CreateToken(param configs.JWT) (t string, err error) {
	ep, _ := utils.ParseDuration(param.ExpiresTime)
	np, _ := utils.ParseDuration(param.BufferTime)
	// CustomClaims因为继承(不是)jwt.RegisteredClaims的结构因此满足Claims接口
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		JwtPayLoad: JwtPayLoad{},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(np)),
			Issuer:    param.Issuer,
		},
	})
	// 自己定义密钥 生成不一样的签名
	str, err := token.SignedString(j.Signed)
	if err != nil {
		return
	}
	return str, err
}
