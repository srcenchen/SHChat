package v1

import (
	"SHChat/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetChatListReq struct {
	g.Meta `method:"get" summary:"获取聊天列表" tags:"聊天"`
}

type GetChatListRes struct {
	ChatList []entity.MessageTable `json:"chatList"`
}
