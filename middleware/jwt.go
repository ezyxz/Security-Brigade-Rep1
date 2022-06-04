package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const signedKey = "Hello World"

type UserClaims struct {
	ID   int64
	time time.Time
	jwt.RegisteredClaims
}

func CreateToken(id int64) string {
	claims := UserClaims{id, time.Now(), jwt.RegisteredClaims{
		NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Hour * time.Duration(1))), // 过期时间4小时
		Issuer:    "FZH",                                                                // 签名的发行者
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(signedKey))
	return tokenString
}

func ParserToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signedKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("无效的token")
			} else {
				return nil, fmt.Errorf("token不可用")
			}
		}
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token无效")
}
