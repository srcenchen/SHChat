package v1

import (
	"SHChat/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetLeastMessageIotReq struct {
	g.Meta `method:"get" summary:"获取最后一个聊天 IOT专用" tags:"聊天"`
}

type GetLeastMessageIotRes struct {
	ChatList []entity.MessageTable `json:"chatList"`
}
