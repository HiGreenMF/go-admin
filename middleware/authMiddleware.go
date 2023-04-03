package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-admin/config"
	"github.com/go-admin/util"
)

// 身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取JWT令牌
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			// 如果JWT令牌不存在，则返回错误响应
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			return
		}

		// 验证JWT令牌
		isValid, err := util.ValidateToken(tokenString)

		if err != nil {
			// 如果验证JWT令牌时发生错误，则返回错误响应
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !isValid {
			// 如果JWT令牌无效，则返回错误响应
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		// 如果JWT令牌有效，则将用户 ID 存储到上下文中
		claims := jwt.MapClaims{}
		token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return []byte(config.Conf.Jwt.SecretKey), nil
		})
		fmt.Println(token)
		fmt.Println(claims)
		userID := claims["user_id"].(string)
		c.Set("userID", string(userID))

		// 继续处理请求
		c.Next()
	}
}
