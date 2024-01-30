package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"douyin/common/ecode"
)

var secretKey = []byte("Q16pjwj5w9Klzs32") // 用于签名和验证JWT的秘密密钥

type UserClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

func CreateTokenByID(userID int64) (string, error) {
	// 创建自定义声明
	claims := UserClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 使用秘密密钥进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseTokenToID(tokenString string) (int64, error) {
	// 解析JWT
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	// 验证JWT是否有效
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, ecode.AuthorizeErr
}
