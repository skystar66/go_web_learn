package module

import "my-hello/app/module/internal"

type User internal.User
//注册请求参数
type UserApiSignUpReq struct {
	Passport string `v:"required|length:6,16#账号不能为空|账号长度应当在min到:max之间"`
	Password string `v:"required|length:6,16#密码不能为空|密码长度应当在min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#请输入确认密码|密码长度应当在min到:max之间|两次输入密码不相同"`
	NickName string `v:"required#昵称不能为空"`
}


//定义modle 结构体
//账号唯一性检测请求,用于前后端交互参数格式约定
type UserApiCheckPassportReq struct {
	Passport string `v:"required#账号不能为空"`
}

//登录请求参数
type UserApiSignInReq struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

//昵称唯一性检测
type UserApiCheckNickNameReq struct {
	NickName string `v:"required#昵称不能为空"`
}

//注册输入参数
type UserServiceSignUpReq struct {
	Passport string
	Password string
	NickName string
}
