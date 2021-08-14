package internal

import "github.com/gogf/gf/os/gtime"

//用户结构体
type User struct {
	Id       uint   `orm:"id,primary" json:"id"`     //用户id
	Passport string `orm:"passport" json:"passport"` //用户账号

	Password string `orm:"password" json:"password"` //用户密码

	NickName string      `orm:"nickName" json:"nickname"`  //用户昵称
	CreateAt *gtime.Time `orm:"create_at" json:"createAt"` //创建时间
	UpdateAt *gtime.Time `orm:"update_at" json:"update_at"` //修改时间
}
