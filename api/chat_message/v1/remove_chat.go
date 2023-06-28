package v1

import "github.com/gogf/gf/v2/frame/g"

type RemoveChatReq struct {
	g.Meta `method:"DELETE" summary:"删除指定聊天记录" tags:"聊天"`
	ChatId string `v:"required#聊天记录id不能为空"`
}

type RemoveChatRes struct {
	IsSuccess bool `json:"isSuccess"`
}
