package service

import (
	"SHChat/internal/model/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var db *gorm.DB

// GetDatabase 获取数据库
func GetDatabase() *gorm.DB {
	once.Do(func() {
		db, _ = gorm.Open(sqlite.Open("./resource/database/data.db"), &gorm.Config{})
	})
	return db
}

// InitData 数据初始化
func InitData() {
	if GetDatabase().AutoMigrate(&entity.UserTable{}, &entity.MessageTable{}) != nil {
		panic("数据库初始化失败")
	}
	// 初始化账号
	var userTable entity.UserTable
	if GetDatabase().Where("id = ?", 1).First(&userTable).RowsAffected == 0 {
		// 没有初始化过账号
		GetDatabase().Create(&entity.UserTable{Password: "sanenchen123", Nickname: "Enchen"})
		GetDatabase().Create(&entity.UserTable{Password: "ye", Nickname: "业"})
	}
}
