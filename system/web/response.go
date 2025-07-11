package web

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var (
	Sucess = Response{Code: 1000, Msg: "success"}
	Error  = Response{Code: 2000, Msg: "error"}

	Failed             = Response{Code: 3000, Msg: "操作失败"}
	FailedUnauthorized = Response{Code: 3001, Msg: "未授权"}
	AccOrPasNotMatch   = Response{Code: 3002, Msg: "账号或密码错误"}
	AccountNotFind     = Response{Code: 3003, Msg: "账号不存在"}
	FailedParam        = Response{Code: 3004, Msg: "参数错误"}
	AccOrPasEmpty      = Response{Code: 3005, Msg: "账号或密码不能为空"}
	AccAlreadyExist    = Response{Code: 3006, Msg: "账号已存在"}
)
