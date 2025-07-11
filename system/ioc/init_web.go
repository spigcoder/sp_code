package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spigcoder/sp_code/system/web"
	"github.com/spigcoder/sp_code/system/web/middleware"
	"strings"
	"time"
)

func InitWeb(handler *web.SysUserHandler) *gin.Engine {
	mds := InitMiddleware()
	// 注册路由
	engine := gin.Default()
	handler.RegisterRouter(engine)
	engine.Use(mds...)
	return engine
}

func InitMiddleware() []gin.HandlerFunc {
	mds := []gin.HandlerFunc{cors.New(cors.Config{
		//这里用来配置允许的域名
		// AllowOrigins:     []string{"https://foo.com"},
		//如果没有这个，那就是默认所有的都可以
		// AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//允许带cookie之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}),
		middleware.NewLoginJWTMiddlewareBuilder().IgnorePaths("/sysuser/login").Build(),
	}
	return mds
}
