package util

import (
	"errors"
	"time"

	"blog-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT 载荷
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(userID uint, username string) (string, error) {
	// 过期时间：24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "blog-backend",
		},
	}

	// 使用 HS256 签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名生成令牌
	tokenString, err := token.SignedString([]byte(config.Load().JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Load().JWTSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	// 验证令牌并提取载荷
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
