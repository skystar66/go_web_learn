package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"my-hello/app/api"
	"my-hello/app/service"
)

func Init() {
	s := g.Server()
	//分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		//中间件执行，拦截等
		group.Middleware(
			service.Middleware.Ctx,
		)
		group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALL("/user/getUserFromSession", api.User.GetUserFromSession)
		})
	})
	s.Run()
}
