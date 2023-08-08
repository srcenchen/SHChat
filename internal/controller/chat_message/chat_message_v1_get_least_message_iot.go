package chat_message

import (
	"SHChat/internal/dao"
	"context"

	"SHChat/api/chat_message/v1"
)

func (c *ControllerV1) GetLeastMessageIot(ctx context.Context, req *v1.GetLeastMessageIotReq) (res *v1.GetLeastMessageIotRes, err error) {
	chatList := dao.GetLeastChat()
	res = &v1.GetLeastMessageIotRes{
		ChatList: chatList,
	}
	return
}
