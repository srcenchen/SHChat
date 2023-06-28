// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package chat_message

import (
	"context"
	
	"SHChat/api/chat_message/v1"
)

type IChatMessageV1 interface {
	ChatListWS(ctx context.Context, req *v1.ChatListWSReq) (res *v1.ChatListWSRes, err error)
	ClearChatList(ctx context.Context, req *v1.ClearChatListReq) (res *v1.ClearChatListRes, err error)
	GetChatList(ctx context.Context, req *v1.GetChatListReq) (res *v1.GetChatListRes, err error)
	RemoveChat(ctx context.Context, req *v1.RemoveChatReq) (res *v1.RemoveChatRes, err error)
	SendChat(ctx context.Context, req *v1.SendChatReq) (res *v1.SendChatRes, err error)
}


