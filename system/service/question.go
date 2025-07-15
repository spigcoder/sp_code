package service

import (
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/repository"
)

type QuestionService interface {
	List(PageNum int, PageSize int, title string, difficulty int32) ([]domain.QuestionVO, int64, error)
}

type QuestionServiceImpl struct {
	repo repository.QestionRepository
}

func NewQuestionServiceImpl(repo repository.QestionRepository) QuestionService {
	return &QuestionServiceImpl{repo: repo}
}

func (impl *QuestionServiceImpl) List(PageNum int, PageSize int, title string, difficulty int32) ([]domain.QuestionVO, int64, error) {
	return impl.repo.List(PageNum, PageSize, title, difficulty)
}
