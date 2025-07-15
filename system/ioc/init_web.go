package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spigcoder/sp_code/system/web"
	"github.com/spigcoder/sp_code/system/web/middleware"
	"strings"
	"time"
)

func InitWeb(SysUserHandler *web.SysUserHandler, QuestionHandler *web.QuestionHandler, rdb redis.Cmdable) *gin.Engine {
	mds := InitMiddleware(rdb)
	// 注册路由
	engine := gin.Default()
	engine.Use(mds...)
	SysUserHandler.RegisterRouter(engine)
	QuestionHandler.RegisterRouter(engine)
	return engine
}

func InitMiddleware(rdb redis.Cmdable) []gin.HandlerFunc {
	mds := []gin.HandlerFunc{cors.New(cors.Config{
		//这里用来配置允许的域名
		AllowOrigins: []string{"https://localhost:5173"},
		//如果没有这个，那就是默认所有的都可以
		//AllowMethods:     []string{"PUT", "PATCH"},
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
		middleware.NewLoginJWTMiddlewareBuilder(rdb).IgnorePaths("/system/user/login").Build(),
	}
	return mds
}
