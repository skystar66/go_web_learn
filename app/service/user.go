package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"my-hello/app/dao"
	"my-hello/app/module"
)

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

//注册用户
func (receiver *serviceUser) Register(user *module.User) error {

	if user.NickName == "" {
		user.NickName = user.Passport
	}

	//账号唯一数据检查
	if res := dao.User.CheckPassport(user.Passport); !res {
		return errors.New(fmt.Sprintf("账号 %s 已经存在",user.Passport))
	}

	//昵称唯一数据检查
	if res := dao.User.CheckNickName(user.NickName); !res {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在",user.NickName))
	}

	dao.User.Save(user)

	return nil
}



//登录用户
func (receiver *serviceUser) Login(r *ghttp.Request,user *module.User) *module.User {
	//账号登录校验
	var userLogin *module.User
	if userLogin=dao.User.Checklogin(user.Passport,user.Password);userLogin==nil{
		return nil
	}
	//存储session
	Session.Set(r.Context(),userLogin)
	return userLogin
}

//获取用户数据从数据库中
func (receiver *serviceUser) GetUserFromDb(passport string) *module.User{
	userDb :=dao.User.GetUserInfo(passport)
	return userDb
}


//获取用户列表
func (receiver *serviceUser) List(r *ghttp.Request) *[]module.User{
	userLists := dao.User.GetUserList()
	return userLists
}

//分页获取用户列表
func (receiver *serviceUser) PageList(page int,limit int) *[]module.User{
	//分页计算公式=(当前页-1)*pageSize
	offset:=(page-1)*limit
	userLists := dao.User.GetUserPageList(offset,limit)
	return userLists
}


