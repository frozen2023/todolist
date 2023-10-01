package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

// MySecret 定义secret
var MySecret = []byte("Frozen")

func MakeToken(id int64) (tokenString string) {
	claim := MyClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, _ = token.SignedString(MySecret)
	return tokenString
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("手写的从前"), nil // 这是我的secret
	}
}

func ParseToken(token string) (any, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &MyClaims{}, Secret())

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	// .(*Type)类型转换
	if claims, ok := parsedToken.Claims.(*MyClaims); ok && parsedToken.Valid {
		// 从Token中解析出id
		return claims.Id, nil
	}
	return nil, errors.New("couldn't handle this token")
}
