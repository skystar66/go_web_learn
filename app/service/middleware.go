package service

import (
	"github.com/gogf/gf/net/ghttp"
	"my-hello/app/module"
	"net/http"
)

//中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

//自定义上下文对象
func (receiver *middlewareService) Ctx(r *ghttp.Request) {
	//初始化
	customCtx := &module.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	//获取用户session登录信息
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &module.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.NickName,
		}
	}
	//执行下一步逻辑
	r.Middleware.Next()

}

//登录鉴权,登陆成功后才能通过
func (receiver *middlewareService) Auth(r *ghttp.Request) {
	if User.IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
