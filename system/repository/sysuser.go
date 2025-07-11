package repository

import (
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/repository/dao"
)

type SysUserRepository interface {
	FindByAccount(account string) (domain.SystemUser, error)
	Add(user domain.SystemUser) error
}

type SysUserServiceImpl struct {
	sysUserDao dao.SysUserDao
}

func NewSysUserRepositoryImpl(sud dao.SysUserDao) SysUserRepository {
	return &SysUserServiceImpl{
		sysUserDao: sud,
	}
}

func (impl *SysUserServiceImpl) FindByAccount(account string) (domain.SystemUser, error) {
	sysUser, err := impl.sysUserDao.FindByAccount(account)
	return impl.convertDaoToDomain(sysUser), err
}

func (impl *SysUserServiceImpl) Add(user domain.SystemUser) error {
	return impl.sysUserDao.Add(impl.convertDomainToDao(user))
}

func (impl *SysUserServiceImpl) convertDomainToDao(user domain.SystemUser) dao.SystemUser {
	return dao.SystemUser{
		Id:       user.Id,
		Account:  user.Account,
		Password: user.Password,
	}
}

func (impl *SysUserServiceImpl) convertDaoToDomain(user dao.SystemUser) domain.SystemUser {
	return domain.SystemUser{
		Id:       user.Id,
		Account:  user.Account,
		Password: user.Password,
	}
}
