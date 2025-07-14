package web

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spigcoder/sp_code/pkg/snowflake"
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/service"
	"github.com/spigcoder/sp_code/system/web/middleware/ijwt"
	"gorm.io/gorm"
	"net/http"
)

type SysUserHandler struct {
	SysUserService service.SysUserService
}

func NewSysUserHandler(sus service.SysUserService) *SysUserHandler {
	return &SysUserHandler{
		SysUserService: sus,
	}
}

func (handler *SysUserHandler) RegisterRouter(server *gin.Engine) {
	group := server.Group("/system/user")
	group.POST("/login", handler.Login)
	group.POST("/add", handler.Add)
	group.POST("/refresh", handler.RefreshJWT)
	group.GET("/info", handler.GetUserInfo)
	group.DELETE("/logout", handler.Logout)
}

func (handler *SysUserHandler) Logout(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, FailedUnauthorized)
		logrus.Error("Get user info failed, 服务器认证失败")
		return
	}
	userClaim, ok := claims.(*ijwt.UserClaims)
	if !ok {
		logrus.Errorf("类型转换失败")
		c.JSON(http.StatusInternalServerError, FailedUnauthorized)
		return
	}
	err := handler.SysUserService.Logout(userClaim.SSID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, FailedLogout)
		return
	}
	c.JSON(http.StatusOK, SucessLogout)
}

func (handler *SysUserHandler) GetUserInfo(c *gin.Context) {
	Claim, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, FailedUnauthorized)
		logrus.Error("Get user info failed, 服务器认证失败")
		return
	}
	userClaim, ok := Claim.(*ijwt.UserClaims)
	if !ok {
		logrus.Errorf("类型转换失败")
		c.JSON(http.StatusInternalServerError, FailedParam)
		return
	}
	NickName, err := handler.SysUserService.GetNickName(userClaim.Uid)
	if err != nil {
		logrus.Errorf("获取用户昵称失败, err: %v", err)
		return
	}
	if err != nil {
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"NickName": NickName})
	return
}

// @title		系统用户接口
// @version	1.0
// @BasePath	/sysuser/add
// @Summary	添加用户
// @Success	200	{string}	string	"ok"
// @Router		/sysuser/add [post]
// @Param		account		query		string	true	"account"
// @Param		password	query		string	true	"password"
// @Failure	500			{string}	string	"internal error"
func (handler *SysUserHandler) Add(c *gin.Context) {
	type Request struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	var req Request
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, FailedParam)
		return
	}
	err := handler.SysUserService.Add(domain.SystemUser{Account: req.Account, Password: req.Password})
	if err == service.AccountAlreadyExist {
		c.JSON(http.StatusInternalServerError, AccAlreadyExist)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, FailedParam)
		logrus.Errorf("未知错误, err: %v", err)
		return
	}
	c.JSON(http.StatusOK, Sucess)
	return
}

func (handler *SysUserHandler) RefreshJWT(c *gin.Context) {
	//获取refreshToken
	refreshToken := c.GetHeader("refresh-token")
	var refreshClaims ijwt.RefreshClaims
	token, err := jwt.ParseWithClaims(refreshToken, &refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return ijwt.RScretKey, nil
	})
	if err != nil || token == nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//这里证明这个长token有效，我们要设置短token
	//设置JWT, 这里同时更新长token和短token
	err = ijwt.SetJWT(c, refreshClaims.Uid, refreshClaims.SSID)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器问题")
		return
	}
}

// @title		系统用户接口
// @version	1.0
// @Summary	注册用户
// @BasePath	/sysuser/login
// @Router		/sysuser/login [post]
// @Param		request	body		domain.SystemUser	true	"request"
// @Success	200		{string}	string				"登录成功"
// @Failure	400		{string}	string				"账号密码不能为空"
// @Failure	500		{string}	string				"用户不存在"
// @Failure	500		{string}	string				"密码或账号错误"
func (handler *SysUserHandler) Login(c *gin.Context) {
	type LoginRequest struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	var request LoginRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, FailedParam)
		return
	}
	if request.Account == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, AccOrPasEmpty)
		return
	}
	sysUser, err := handler.SysUserService.Login(domain.SystemUser{Account: request.Account, Password: request.Password})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, AccountNotFind)
			return
		} else if err == service.PasswordNotMatch {
			c.JSON(http.StatusInternalServerError, AccOrPasNotMatch)
			return
		}
	}
	jwtId := snowflake.GenID()
	err = ijwt.SetJWT(c, sysUser.Id, jwtId)
	handler.SysUserService.SetJwtValid(jwtId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, FailedParam)
		return
	}
	c.JSON(http.StatusOK, Sucess)
	return
}
