package domain

import "time"

type SystemUser struct {
	Id        int64
	Account   string
	Password  string
	NickName  string
	CreatedAt time.Time
}
