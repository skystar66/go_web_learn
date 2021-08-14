package dao

import (
	"github.com/gogf/gf/frame/g"
	"my-hello/app/module"
)

//user数据持久化服务
var User = userDao{}

type userDao struct {
}

//注册用户
func (receiver userDao) Save(user *module.User) {
	g.DB().Model("user").Data(user).Insert()
}

// 账号唯一性数据检查,true 表示唯一，false 不唯一
func (receiver *userDao) CheckPassport(passport string) bool {
	if count, err := g.DB().Model("user").Where("passport", passport).Count(); err != nil {
		return count == 0
	}
	return false
}

// 昵称唯一性数据检查,true 表示唯一，false 不唯一
func (receiver *userDao) CheckNickName(nickName string) bool {
	if count, err := g.DB().Model("user").Where("nickName", nickName).Count(); err != nil {
		return count == 0
	}
	return false
}

//获取用户详细信息
func (receiver userDao) GetUserInfo(nickName string) *module.User {
	var user = &module.User{}
	g.DB().Model("user").Where("nickName", nickName).Scan(user)
	return user
}

//获取所有用户数据
func (receiver userDao) GetUserList() *[]module.User {
	var users = &[]module.User{}
	g.DB().Model("user").Scan(users)
	return users
}