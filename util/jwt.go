package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-admin/config"
)

// JWTClaims JWT 负载数据
type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// CreateToken 创建 JWT Token
func CreateToken(userID string) (string, error) {
	// 创建一个 Token 对象
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := JWTClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "super_admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用秘钥进行签名，并返回签名后的 Token 字符串
	return token.SignedString([]byte(config.Conf.Jwt.SecretKey))
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string) (*JWTClaims, error) {
	// 解析 Token 字符串，并验证签名
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.Conf.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// 检查 Token 是否过期
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// 返回解析后的负载数据
	return claims, nil
}

// 验证JWT令牌
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// 验证JWT令牌的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回JWT密钥
		return []byte(config.Conf.Jwt.SecretKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
