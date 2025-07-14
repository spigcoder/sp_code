package dao

import "time"

type Question struct {
	Id           int64
	Title        string    `gorm:"not null;type:varchar(255)"`
	Difficulty   int32     `gorm:"not null"`
	Content      string    `gorm:"not null;type:varchar(1000)"`
	TimeLimit    int32     `gorm:"not null"`
	SpaceLimit   int32     `gorm:"not null"`
	QuestionCase string    `gorm:"not null;type:varchar(500)"`
	DefaultCode  string    `gorm:"not null;type:varchar(500)"`
	MainCode     string    `gorm:"not null;type:varchar(500)"`
	CreatedAt    time.Time `gorm:"not null"`
	CreatedBy    int64     `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
	DeletedAt    time.Time `gorm:"not null"`
}
