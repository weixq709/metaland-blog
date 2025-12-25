package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
)

// TODO 改为动态文件
const secret = "123456"

// Generate 根据用户名生成token字符串
func Generate(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userName,
		"exp":      time.Now().Add(30 * time.Minute).Unix(),
	})
	return token.SignedString([]byte(secret))
}

// Parse 根据token解析用户信息，如果解析token失败或token非法，抛出异常
// 如果解析成功，将用户信息以map返回
func Parse(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if token == nil || err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid expired")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Wrap(err, "failed to parse token")
	}
	return claims, nil
}
