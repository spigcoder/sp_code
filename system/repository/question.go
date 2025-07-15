package repository

import (
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/repository/dao"
)

type QestionRepository interface {
	List(PageNum int, PageSize int, title string, difficulty int32) ([]domain.QuestionVO, int64, error)
}

type QestionRepositoryImpl struct {
	questionDao dao.QuestionDao
}

func NewQuestionRepositoryImpl(questionDao dao.QuestionDao) QestionRepository {
	return &QestionRepositoryImpl{
		questionDao: questionDao,
	}
}

func (impl *QestionRepositoryImpl) List(PageNum int, PageSize int, title string, difficulty int32) ([]domain.QuestionVO, int64, error) {
	return impl.questionDao.List(PageNum, PageSize, title, difficulty)
}
