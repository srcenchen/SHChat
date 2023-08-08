package cmd

import (
	"SHChat/internal/controller/chat_message"
	"SHChat/internal/controller/user"
	"SHChat/internal/service"
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "SHChat",
		Usage: "SHChat",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initDir()          // 初始化文件夹
			service.InitData() // 初始化数据
			s := g.Server()

			// 静态资源绑定
			s.SetIndexFolder(true)
			s.SetServerRoot("public")
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				r.Response.ServeFile("public/index.html")
			})

			api := s.Group("/api")
			api.Group("/user", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.NewV1(),
				)
			})
			api.Group("/message", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					chat_message.NewV1(),
				)
			})
			s.EnableHTTPS("manifest/certs/shchat.pem", "manifest/certs/shchat.key")
			s.Run()
			return nil
		},
	}
)

// initDir /** 初始化文件夹
func initDir() {
	// 创建文件夹
	createDir("./resource")
	createDir("./resource/database")
}

// createDir /** 创建文件夹
func createDir(path string) {
	// 检测是否存在 path 文件夹 如果不存在则创建
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
}
