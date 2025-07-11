package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spigcoder/sp_code/pkg/snowflake"
	_ "github.com/spigcoder/sp_code/system/docs"
	"github.com/spigcoder/sp_code/system/ioc"
	"github.com/spigcoder/sp_code/system/startup"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	snowflake.Init(10)
	ioc.InitLogrus()
	engine := startup.InitWebServer()
	go func() {
		gin.Default().GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}()
	engine.Run(":8080")
}
