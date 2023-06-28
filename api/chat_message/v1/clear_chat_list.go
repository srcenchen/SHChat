package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ClearChatListReq struct {
	g.Meta `method:"DELETE" summary:"删除聊天记录" tags:"聊天"`
}

type ClearChatListRes struct {
	IsSuccess bool `json:"isSuccess"`
}
