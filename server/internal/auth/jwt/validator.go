package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func (j *JWT) ParseToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Signed, nil // 此处返回参数interface{}是自己的密钥
	})
	if err != nil {
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				return nil, TokenExpired
			case errors.Is(err, jwt.ErrTokenMalformed):
				return nil, TokenMalformed
			case errors.Is(err, jwt.ErrTokenSignatureInvalid):
				return nil, TokenSignatureInvalid
			case errors.Is(err, jwt.ErrTokenNotValidYet):
				return nil, TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		}
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return nil, err
}
