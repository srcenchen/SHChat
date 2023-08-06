package chat_message

import (
	"SHChat/internal/dao"
	"SHChat/internal/logic/chat_message"
	"context"

	"SHChat/api/chat_message/v1"
)

func (c *ControllerV1) ClearChatList(_ context.Context, _ *v1.ClearChatListReq) (res *v1.ClearChatListRes, err error) {
	dao.ClearChat()
	chat_message.SendBroadCast("", "")
	res = &v1.ClearChatListRes{
		IsSuccess: true,
	}
	return
}
