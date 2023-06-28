package v1

import "github.com/gogf/gf/v2/frame/g"

type SendChatReq struct {
	g.Meta   `method:"PUT" summary:"发送聊天" tags:"聊天"`
	Message  string `v:"required#聊天内容不能为空"`
	NickName string `v:"required#昵称不能为空"`
}

type SendChatRes struct {
	IsSuccess bool `json:"isSuccess"`
}
