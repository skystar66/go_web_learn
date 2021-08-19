package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {

}
//普通struct转换
func Struct_conv() {


	type User struct {
		Uid  int
		Name string
	}
	params := g.Map{
		"uid":  1,
		"name": "john",
	}
	var user *User
	if err := gconv.Struct(params, &user); err != nil {
		panic(err)
	}
	g.Dump(user)

}
//递归转换
func struct_deep_conv() {
	type Ids struct {
		Id         int    `json:"id"`
		Uid        int    `json:"uid"`
	}
	type Base struct {
		Ids
		CreateTime string `json:"create_time"`
	}
	type User struct {
		Base
		Passport   string `json:"passport"`
		Password   string `json:"password"`
		Nickname   string `json:"nickname"`
	}
	data := g.Map{
		"id"          : 1,
		"uid"         : 100,
		"passport"    : "john",
		"password"    : "123456",
		"nickname"    : "John",
		"create_time" : "2019",
	}
	user := new(User)
	gconv.Struct(data, user)
	g.Dump(user)
}
//实例一 基本使用
func first() {
	type User struct {
		Uid      int
		Name     string
		SiteUrl  string
		NickName string
		Pass1    string `c:"password1"`
		Pass2    string `c:"password2"`
	}

	var user *User

	// 使用默认映射规则绑定属性值到对象
	user = new(User)
	params1 := g.Map{
		"uid":       1,
		"Name":      "john",
		"site_url":  "https://goframe.org",
		"nick_name": "johng",
		"PASS1":     "123",
		"PASS2":     "456",
	}
	if err := gconv.Struct(params1, user); err == nil {
		g.Dump(user)
	}

	// 使用struct tag映射绑定属性值到对象
	user = new(User)
	params2 := g.Map{
		"uid":       2,
		"name":      "smith",
		"site-url":  "https://goframe.org",
		"nick name": "johng",
		"password1": "111",
		"password2": "222",
	}
	if err := gconv.Struct(params2, user); err == nil {
		g.Dump(user)
	}
	/*	可以看到，我们可以直接通过Struct方法将map按照默认规则绑定到struct上，
	也可以使用struct tag的方式进行灵活的设置。此外，Struct方法有第三个map参数，用于指定自定义的参数名称到属性名称的映射关系。
	*/
}

//实例二，复杂使用

func sencond() {
	//	属性支持struct对象或者struct对象指针（目标为指针且未nil时，转换时会自动初始化）转换。

	type Score struct {
		Name   string
		Result int
	}
	type User1 struct {
		Scores Score
	}
	type User2 struct {
		Scores *Score
	}

	user1  := new(User1)
	user2  := new(User2)
	scores := g.Map{
		"Scores": g.Map{
			"Name":   "john",
			"Result": 100,
		},
	}

	if err := gconv.Struct(scores, user1); err != nil {
		fmt.Println(err)
	} else {
		g.Dump(user1)
	}
	if err := gconv.Struct(scores, user2); err != nil {
		fmt.Println(err)
	} else {
		g.Dump(user2)
	}


}

