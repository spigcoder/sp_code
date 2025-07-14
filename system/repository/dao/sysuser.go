package dao

import (
	"gorm.io/gorm"
	"time"
)

type SystemUser struct {
	Id        int64  `gorm:"primaryKey;autoIncrement:false"`
	Account   string `gorm:"type:varchar(54);not null;uniqueIndex:account_password"`
	Password  string `gorm:"type:varchar(108);not null;uniqueIndex:account_password"`
	NickName  string `gorm:"type:varchar(54);not null"`
	CreatedAt time.Time
}

type SysUserDao interface {
	FindByAccount(account string) (SystemUser, error)
	Add(user SystemUser) error
	GetUserById(userId int64) (SystemUser, error)
}

type SysUserDaoImpl struct {
	db *gorm.DB
}

func NewSysUserDaoImpl(db *gorm.DB) SysUserDao {
	return &SysUserDaoImpl{
		db: db,
	}
}

func (impl *SysUserDaoImpl) GetUserById(userId int64) (SystemUser, error) {
	var user SystemUser
	err := impl.db.Where("id =?", userId).First(&user).Error
	return user, err
}

func (impl *SysUserDaoImpl) FindByAccount(account string) (SystemUser, error) {
	var user SystemUser
	err := impl.db.Where("account = ?", account).First(&user).Error
	return user, err
}

func (impl *SysUserDaoImpl) Add(user SystemUser) error {
	return impl.db.Create(&user).Error
}
