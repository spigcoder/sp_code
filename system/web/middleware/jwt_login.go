package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spigcoder/sp_code/system/web/middleware/ijwt"
	"net/http"
)

type LoginJWTMiddlewareBuilder struct {
	path []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.path = append(l.path, path)
	return l
}

// 如果这里由退出登录这个功能的话，我们就需要在检查token之前，检查他是否在Reids的退出token种
// 如果在的话，那么证明已经退出登录了，就返回登录失败
func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, path := range l.path {
			if path == c.Request.URL.Path {
				return
			}
		}
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userClaims := &ijwt.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenHeader, userClaims, func(token *jwt.Token) (interface{}, error) {
			return ijwt.AScretKey, nil
		})
		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 检查userAgent是否一致, 防止token劫持
		if userClaims.UserAgent != c.Request.UserAgent() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("claims", userClaims)
	}
}
