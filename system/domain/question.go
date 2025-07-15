package domain

import "time"

type Question struct {
	Id           int64
	Title        string
	Difficulty   int32
	TimeLimit    int64
	SpaceLimit   int64
	Content      string
	QuestionCase string
	DefaultCode  string
	MainCode     string
	CreateAt     time.Time
	UpdateAt     time.Time
	CreateBy     int64
}

type QuestionVO struct {
	ID         int64
	Title      string
	Difficulty int32
	CreateName string
	CreateAt   time.Time
}
