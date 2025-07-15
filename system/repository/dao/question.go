package dao

import (
	"github.com/spigcoder/sp_code/system/domain"
	"gorm.io/gorm"
	"time"
)

type Question struct {
	Id           int64
	Title        string    `gorm:"not null;type:varchar(255);uniqueIndex:difficulty_title"`
	Difficulty   int32     `gorm:"not null;uniqueIndex:difficulty_title"`
	Content      string    `gorm:"not null;type:varchar(1000)"`
	TimeLimit    int64     `gorm:"not null"`
	SpaceLimit   int64     `gorm:"not null"`
	QuestionCase string    `gorm:"not null;type:varchar(1000)"`
	DefaultCode  string    `gorm:"not null;type:varchar(1000)"`
	MainCode     string    `gorm:"not null;type:varchar(1000)"`
	CreatedAt    time.Time `gorm:"not null"`
	CreatedBy    int64     `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

type QuestionDao interface {
	List(pageNum int, pageSize int, title string, difficulty int32) ([]domain.QuestionVO, int64, error)
}

type QuestionDaoImpl struct {
	db *gorm.DB
}

func NewQuestionDaoImpl(db *gorm.DB) QuestionDao {
	return &QuestionDaoImpl{db: db}
}

func (impl *QuestionDaoImpl) List(PageNum int, PageSize int, title string, difficulty int32) ([]domain.QuestionVO, int64, error) {
	var (
		results []domain.QuestionVO
		total   int64
	)
	baseQuery := impl.db.Model(&Question{})
	if difficulty > 0 && difficulty < 4 {
		baseQuery = baseQuery.Where("questions.difficulty = ?", difficulty)
	}
	if title != "" {
		baseQuery = baseQuery.Where("questions.title LIKE ?", title+"%")
	}
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	query := baseQuery.
		Select(`
			questions.id,
			questions.title,
			questions.difficulty,
			system_users.nick_name AS create_name,
			questions.created_at
		`).Joins("LEFT JOIN system_users ON questions.created_by = system_users.id").
		Order("questions.created_at DESC")
	if PageNum > 0 && PageSize > 0 {
		offset := PageSize * (PageNum - 1)
		query = query.Offset(offset).Limit(PageSize)
	}
	err := query.Scan(&results).Error
	return results, total, err
}
