package push_service

import (
	"fmt"
	"github.com/Scorpio69t/jpush-api-golang-client"
)

func Push(alert string) {
	var pf jpush.Platform
	_ = pf.Add(jpush.ANDROID)
	var at jpush.Audience
	at.SetAlias([]string{"shchat"})
	// Notification
	var n jpush.Notification
	n.Android = &jpush.AndroidNotification{Alert: alert, Title: "SHChat 消息通知", Priority: 2}
	// PayLoad
	payload := jpush.NewPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&at)
	payload.SetNotification(&n)
	// Send
	c := jpush.NewJPushClient("79558f0006d7136b7fce7360", "532729df6cc3d27e2f1e49b1") // appKey and masterSecret can be gotten from https://www.jiguang.cn/
	data, err := payload.Bytes()
	if err != nil {
		panic(err)
	}
	res, err := c.Push(data)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("ok: %v\n", res)
	}
}
