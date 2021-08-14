package module

import "github.com/gogf/gf/net/ghttp"

const (
	//上下文变量存储的健名
	ContextKey = "ContextKey"
)

type Context struct {
	Session *ghttp.Session //当前session管理对象
	User    *ContextUser   //上下文用户信息
}

type ContextUser struct {
	Id       uint   //用户id
	Passport string //用户账号
	Nickname string //用户名称
}
