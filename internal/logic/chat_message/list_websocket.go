package chat_message

import (
	"SHChat/internal/dao"
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	wsInfo = gmap.New(true) // 使用默认的并发安全Map
)

func ChatListWS(ctx context.Context) {
	ws, _ := g.RequestFromCtx(ctx).WebSocket()
	wsInfo.Set(ws, ws)
	for {
		// 先发送一次
		_ = ws.WriteJSON(
			g.Map{
				"code": 0,
				"msg":  "首次连接推送",
				"data": g.Map{
					"chatList": dao.GetChatList(),
				},
			},
		)
		// 堵塞读取
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func SendBroadCast() {
	// 连接池广播
	wsInfo.RLockFunc(func(m map[interface{}]interface{}) {
		for ws := range m {
			_ = m[ws].(*ghttp.WebSocket).WriteJSON(
				g.Map{
					"code": 0,
					"msg":  "消息广播",
					"data": g.Map{
						"chatList": dao.GetChatList(),
					},
				})
		}
	})
}
