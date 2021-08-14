package service

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"my-hello/app/module"
)

var Context = contextService{}

type contextService struct{}

//上下文初始化,初始化上下文对象指针到上下文对象中
func (receiver *contextService) Init(r *ghttp.Request, customCtx *module.Context) {
	r.SetCtxVar(module.ContextKey, customCtx)
}

//获得上下文变量，如果没有设置 返回nil
func (receiver *contextService) Get(ctx context.Context) *module.Context {
	value := ctx.Value(module.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*module.Context); ok {
		return localCtx
	}
	return nil
}

//将上下文信息完全设置到请求的上下文中，完全父覆盖
func (receiver *contextService) Set(ctx context.Context, contextUser *module.ContextUser) {
	receiver.Get(ctx).User = contextUser
}
