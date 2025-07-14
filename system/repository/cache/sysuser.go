package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type SysUserCache interface {
	GetNickName(userId int64) (string, error)
	SetNickName(userId int64, nickName string, ttl time.Duration) error
	DeleteJwt(JwtId int64) error
	SetJwtInvalid(userId int64) error
	SetJwtValid(userId int64) error
}

const (
	KeyUserNickName = "user:id:"
	KeyJwt          = "jwt:valid:"
)

type SysUserCacheImpl struct {
	rdb redis.Cmdable
}

func NewSysUserCacheImpl(rdb redis.Cmdable) SysUserCache {
	return &SysUserCacheImpl{
		rdb: rdb,
	}
}

func (s *SysUserCacheImpl) SetJwtValid(jwtId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return s.rdb.Set(ctx, GetJwtKey(jwtId), true, time.Hour*2).Err()
}

func (s *SysUserCacheImpl) SetJwtInvalid(jwtId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return s.rdb.Set(ctx, GetJwtKey(jwtId), false, time.Hour*2).Err()
}

func (s *SysUserCacheImpl) DeleteJwt(jwtId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return s.rdb.Del(ctx, GetJwtKey(jwtId)).Err()
}

func (s *SysUserCacheImpl) GetNickName(userId int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return s.rdb.Get(ctx, GetNickKey(userId)).Result()
}

func (s *SysUserCacheImpl) SetNickName(userId int64, nickName string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return s.rdb.Set(ctx, GetNickKey(userId), nickName, ttl).Err()
}

func GetJwtKey(userId int64) string {
	return KeyJwt + strconv.FormatInt(userId, 10)
}

func GetNickKey(userId int64) string {
	return KeyUserNickName + strconv.FormatInt(userId, 10)
}
