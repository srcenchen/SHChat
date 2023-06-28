package chat_message

import (
	"SHChat/internal/dao"
	"context"

	"SHChat/api/chat_message/v1"
)

func (c *ControllerV1) GetChatList(_ context.Context, req *v1.GetChatListReq) (res *v1.GetChatListRes, err error) {
	chatList := dao.GetChatList()
	res = &v1.GetChatListRes{
		ChatList: chatList,
	}
	return
}
