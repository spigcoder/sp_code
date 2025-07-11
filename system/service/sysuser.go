package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spigcoder/sp_code/pkg/snowflake"
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/repository"
	"github.com/spigcoder/sp_code/system/utils/bcrypt"
	"gorm.io/gorm"
)

var (
	PasswordNotMatch    = errors.New("密码错误")
	AccountAlreadyExist = errors.New("账号已存在")
)

type SysUserService interface {
	Login(user domain.SystemUser) (domain.SystemUser, error)
	Add(user domain.SystemUser) error
}

type SysUserServiceImpl struct {
	SysUserRepo repository.SysUserRepository
}

func NewSysUserServiceImpl(sur repository.SysUserRepository) SysUserService {
	return &SysUserServiceImpl{
		SysUserRepo: sur,
	}
}

func (impl *SysUserServiceImpl) Add(sysUser domain.SystemUser) error {
	password, err := bcrypt.Encrypt(sysUser.Password)
	if err != nil {
		logrus.Errorf("加密密码失败, err: %v", err)
		return err
	}
	sysUser.Password = password
	sysUser.Id = snowflake.GenID()
	err = impl.SysUserRepo.Add(sysUser)
	if err == gorm.ErrDuplicatedKey {
		return PasswordNotMatch
	}
	return err
}

func (impl *SysUserServiceImpl) Login(user domain.SystemUser) (domain.SystemUser, error) {
	sysUser, err := impl.SysUserRepo.FindByAccount(user.Account)
	if err != nil {
		return domain.SystemUser{}, err
	}
	if !bcrypt.CompareHashAndPassword(sysUser.Password, user.Password) {
		return domain.SystemUser{}, PasswordNotMatch
	}
	return sysUser, nil
}
