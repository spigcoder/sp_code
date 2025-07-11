package domain

import "time"

type SystemUser struct {
	Id        int64
	Account   string
	Password  string
	CreatedAt time.Time
}
