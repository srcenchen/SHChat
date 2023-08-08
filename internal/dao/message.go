package dao

import (
	"SHChat/internal/model/entity"
	"SHChat/internal/service"
	"time"
)

// GetChatList 获取聊天列表
func GetChatList() []entity.MessageTable {
	var chatList []entity.MessageTable
	service.GetDatabase().Find(&chatList)
	return chatList
}

// GetLeastChat 获取最新一条聊天记录
func GetLeastChat() []entity.MessageTable {
	var chat []entity.MessageTable
	service.GetDatabase().Last(&chat)
	return chat
}

// SendChat 添加聊天记录
func SendChat(message string, nickname string) {
	// 获取当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	service.GetDatabase().Create(&entity.MessageTable{Content: message, NickName: nickname, Date: nowTime})
}

// RemoveChat 删除聊天记录
func RemoveChat(id string) {
	service.GetDatabase().Where("id = ?", id).Delete(&entity.MessageTable{})
}

// ClearChat 清空聊天记录
func ClearChat() {
	service.GetDatabase().Where("id != -1").Delete(&entity.MessageTable{})
}
