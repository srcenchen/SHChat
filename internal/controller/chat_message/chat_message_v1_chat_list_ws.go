package chat_message

import (
	"SHChat/internal/logic/chat_message"
	"context"

	"SHChat/api/chat_message/v1"
)

func (c *ControllerV1) ChatListWS(ctx context.Context, _ *v1.ChatListWSReq) (res *v1.ChatListWSRes, err error) {
	chat_message.ChatListWS(ctx)
	return &v1.ChatListWSRes{}, nil
}
