package chat_message

import (
	"SHChat/internal/dao"
	"SHChat/internal/logic/chat_message"
	"context"

	"SHChat/api/chat_message/v1"
)

func (c *ControllerV1) RemoveChat(_ context.Context, req *v1.RemoveChatReq) (res *v1.RemoveChatRes, err error) {
	dao.RemoveChat(req.ChatId)
	chat_message.SendBroadCast()
	res = &v1.RemoveChatRes{
		IsSuccess: true,
	}
	return
}
