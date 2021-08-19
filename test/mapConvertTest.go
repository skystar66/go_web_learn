package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"reflect"
)

func main() {

	type User struct {
		Uid  int    `c:"uid"`
		Name string `c:"name"`
	}
	// 对象
	g.Dump(gconv.Map(User{
		Uid:  1,
		Name: "john",
	}))
	// 对象指针
	g.Dump(gconv.Map(&User{
		Uid:  1,
		Name: "john",
	}))

	// 任意map类型
	g.Dump(gconv.Map(map[int]int{
		100: 10000,
	}))


	DeepDump()
}



func DeepDump() {
	type Base struct {
		Id   int    `c:"id"`
		Date string `c:"date"`
	}
	type User struct {
		UserBase Base   `c:"base"`
		Passport string `c:"passport"`
		Password string `c:"password"`
		Nickname string `c:"nickname"`
	}
	user := User{
		UserBase: Base{
			Id:   1,
			Date: "2019-10-01",
		},
		Passport: "john",
		Password: "123456",
		Nickname: "JohnGuo",
	}
	m1 := gconv.Map(user)
	m2 := gconv.MapDeep(user)
	g.Dump(m1, m2)
	fmt.Println(reflect.TypeOf(m1["base"]))
	fmt.Println(reflect.TypeOf(m2["base"]))
}










