package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/service"
	"github.com/spigcoder/sp_code/system/web/response"
	"net/http"
)

type QuestionHandler struct {
	service service.QuestionService
}

func NewQuestionHandler(service service.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: service}
}

func (handler *QuestionHandler) RegisterRouter(server *gin.Engine) {
	group := server.Group("/system/question")
	group.GET("/list", handler.List)
}

func (handler *QuestionHandler) List(c *gin.Context) {
	type Request struct {
		PageNum    int    `form:"pageNum"`
		PageSize   int    `form:"pageSize"`
		Title      string `form:"title"`
		Difficulty int32  `form:"difficulty"`
	}
	request := Request{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Failed)
		logrus.Errorf("bind request failed, err:%v", err)
		return
	}
	if request.PageNum == 0 {
		request.PageNum = 1
	}
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	questions, total, err := handler.service.List(request.PageNum, request.PageSize, request.Title, request.Difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Failed)
		logrus.Errorf("list questions failed, err:%v", err)
		return
	}
	c.JSON(http.StatusOK, response.ListData[domain.QuestionVO]{
		Rows:  questions,
		Total: total,
		Code:  http.StatusOK,
		Msg:   "success",
	})
	return
}
