package dao

import (
	"SHChat/internal/model/entity"
	"SHChat/internal/service"
)

// GetUserInfoById 获取用户信息
func GetUserInfoById(id string) entity.UserTable {
	var userTable entity.UserTable
	service.GetDatabase().Where("id = ?", id).First(&userTable)
	return userTable
}

// GetUserInfoByPasswd 获取用户信息
func GetUserInfoByPasswd(passwd string) int64 {
	var userTable entity.UserTable
	return service.GetDatabase().Where(entity.UserTable{Password: passwd}).First(&userTable).RowsAffected
}
