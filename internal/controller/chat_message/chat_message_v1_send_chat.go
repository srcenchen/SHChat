package chat_message

import (
	"SHChat/api/chat_message/v1"
	"SHChat/internal/dao"
	"SHChat/internal/logic/chat_message"
	"context"
)

func (c *ControllerV1) SendChat(_ context.Context, req *v1.SendChatReq) (res *v1.SendChatRes, err error) {
	dao.SendChat(req.Message, req.NickName)
	chat_message.SendBroadCast()
	res = &v1.SendChatRes{
		IsSuccess: true,
	}
	return
}
