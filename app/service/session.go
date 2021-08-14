package service

import (
	"context"
	"my-hello/app/module"
)

//Session管理服务
var Session = sessionService{}

type sessionService struct{}

const (
	sessionKey = "sessionKey"
)

//设置用户session
func (receiver *sessionService) Set(ctx context.Context, user *module.User) error {
	return Context.Get(ctx).Session.Set(sessionKey, user)
}

//获取当前登录用户信息，如果用户未登录，返回nil
func (receiver *sessionService) GetUser(ctx context.Context) *module.User {

	customCtx := Context.Get(ctx)
	if customCtx == nil {
		return nil
	}
	if localUser := customCtx.Session.GetVar(sessionKey); localUser != nil {
		var user *module.User
		//字段映射
		localUser.Struct(&user)
		return user
	}
	return nil
}

//删除用户session
func (receiver *sessionService) RemoveUserSession(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKey)
	}
	return nil
}
