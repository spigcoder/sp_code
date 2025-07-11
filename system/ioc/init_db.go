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
	user := dao.SystemUser{
		Id:       1943522465699373056,
		Account:  "sb",
		Password: "$2a$10$BuCBsohLlvvpFJpHkWC0MOxKUZ7b.w4ky712xWoM3HYcvt7aVT5GS",
	}
	db.Create(&user)
	return db
}
