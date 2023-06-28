package v1

import "github.com/gogf/gf/v2/frame/g"

type ChatListWSReq struct {
	g.Meta `method:"GET" summary:"获取信息 ws" tags:"聊天Ws"`
}

type ChatListWSRes struct {
}
