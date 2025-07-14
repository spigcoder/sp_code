package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/spigcoder/sp_code/system/domain"
	"github.com/spigcoder/sp_code/system/repository/cache"
	"github.com/spigcoder/sp_code/system/repository/dao"
	"time"
)

type SysUserRepository interface {
	FindByAccount(account string) (domain.SystemUser, error)
	Add(user domain.SystemUser) error
	GetUserById(userId int64) (domain.SystemUser, error)
	GetNickName(userId int64) (string, error)
	SetNickName(userId int64, nickName string) error
	DeleteJwt(jwtId int64) error
	SetJwtInvalid(jwtId int64) error
	SetJwtValid(jwtId int64) error
}

type SysUserServiceImpl struct {
	sysUserDao   dao.SysUserDao
	sysUserCache cache.SysUserCache
}

func NewSysUserRepositoryImpl(sud dao.SysUserDao, cache cache.SysUserCache) SysUserRepository {
	return &SysUserServiceImpl{
		sysUserDao:   sud,
		sysUserCache: cache,
	}
}

func (impl *SysUserServiceImpl) SetJwtValid(jwtId int64) error {
	return impl.sysUserCache.SetJwtValid(jwtId)
}

func (impl *SysUserServiceImpl) SetJwtInvalid(jwtId int64) error {
	return impl.sysUserCache.SetJwtInvalid(jwtId)
}

func (impl *SysUserServiceImpl) DeleteJwt(JwtId int64) error {
	return impl.sysUserCache.DeleteJwt(JwtId)
}

func (impl *SysUserServiceImpl) SetNickName(userId int64, nickName string) error {
	return impl.sysUserCache.SetNickName(userId, nickName, time.Minute*2)
}

func (impl *SysUserServiceImpl) GetNickName(userId int64) (string, error) {
	nickName, err := impl.sysUserCache.GetNickName(userId)
	if err != nil {
		sysUser, err := impl.sysUserDao.GetUserById(userId)
		if err != nil {
			logrus.Errorf("获取用户信息失败, err: %v", err)
			return "", err
		}
		impl.SetNickName(userId, sysUser.NickName)
		return sysUser.NickName, nil
	}
	return nickName, nil
}

func (impl *SysUserServiceImpl) GetUserById(userId int64) (domain.SystemUser, error) {
	sysUser, err := impl.sysUserDao.GetUserById(userId)
	if err != nil {
		return domain.SystemUser{}, err
	}
	impl.SetNickName(userId, sysUser.NickName)
	return impl.convertDaoToDomain(sysUser), err
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
		NickName: user.NickName,
	}
}

func (impl *SysUserServiceImpl) convertDaoToDomain(user dao.SystemUser) domain.SystemUser {
	return domain.SystemUser{
		Id:       user.Id,
		Account:  user.Account,
		Password: user.Password,
		NickName: user.NickName,
	}
}
