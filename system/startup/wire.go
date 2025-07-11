//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spigcoder/sp_code/system/ioc"
	"github.com/spigcoder/sp_code/system/repository"
	"github.com/spigcoder/sp_code/system/repository/dao"
	"github.com/spigcoder/sp_code/system/service"
	"github.com/spigcoder/sp_code/system/web"
)

// 定义 Provider Set
var WebServerSet = wire.NewSet(
	ioc.InitDB,                          // 提供 *gorm.DB
	dao.NewSysUserDaoImpl,               // 依赖 *gorm.DB
	repository.NewSysUserRepositoryImpl, // 依赖 dao.SysUserDao
	service.NewSysUserServiceImpl,       // 依赖 repository.SysUserRepository
	web.NewSysUserHandler,               // 依赖 service.SysUserService
	ioc.InitWeb,                         // 依赖 web.SysUserHandler, 返回 *gin.Engine
)

func InitWebServer() *gin.Engine {
	wire.Build(WebServerSet)
	return nil
}
