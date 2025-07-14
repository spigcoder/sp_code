package ioc

import (
	"github.com/spigcoder/sp_code/system/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "spigcoder:123456@tcp(127.0.0.1:3306)/sp_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dao.InitTable(db)
	return db
}
