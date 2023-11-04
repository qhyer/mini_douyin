package jwt

import (
	"douyin/common/ecode"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
}

var secretKey = []byte("Q16pjwj5w9Klzs32") // 用于签名和验证JWT的秘密密钥

func (*JWT) CreateTokenByID(userID int64) (string, error) {
	// 创建自定义声明
	claims := jwt.MapClaims{
		"UserID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 过期时间为24小时后
	}

	// 使用秘密密钥进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (*JWT) ParseTokenToID(tokenString string) (int64, error) {
	// 解析JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	// 验证JWT是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["UserID"].(int64); ok {
			return userID, nil
		}
	}

	return 0, ecode.AuthorizeErr
}
