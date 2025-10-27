package middleware

import (
	"net/http"
	"strings"

	"blog-backend/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthRequired 验证 JWT 的中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logrus.Warn("请求头缺少 Authorization")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查 Bearer 前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			logrus.Warn("Authorization 格式错误")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// 解析令牌
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			logrus.Warnf("令牌解析失败: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
