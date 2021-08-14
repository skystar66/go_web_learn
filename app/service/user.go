package service

import "context"

//User中间件服务管理
var User = serviceUser{}

type serviceUser struct {
}

//判断用户是否已经登录
func (receiver *serviceUser) IsSignedIn(ctx context.Context) bool {
	customCtx := Context.Get(ctx)
	if customCtx == nil {
		return false
	}
	if customCtx.User != nil {
		return true
	}
	return false
}
