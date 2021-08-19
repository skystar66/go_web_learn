package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {

	scan_strut()
	scan_struts()
	scan_map()
	scan_maps()

}
//自动识别转换Struct
func scan_strut() {
	type User struct {
		Uid  int
		Name string
	}
	params := g.Map{
		"uid":  1,
		"name": "john",
	}
	var user *User
	if err := gconv.Scan(params, &user); err != nil {
		panic(err)
	}
	g.Dump(user)
}
//自动识别转换Struct数组
func scan_struts() {
	type User struct {
		Uid  int
		Name string
	}
	params := g.Slice{
		g.Map{
			"uid":  1,
			"name": "john",
		},
		g.Map{
			"uid":  2,
			"name": "smith",
		},
	}
	var users []*User
	if err := gconv.Scan(params, &users); err != nil {
		panic(err)
	}
	g.Dump(users)
}

//自动识别转换map
func scan_map() {
	var (
		user   map[string]string
		params = g.Map{
			"uid":  1,
			"name": "john",
		}
	)
	if err := gconv.Scan(params, &user); err != nil {
		panic(err)
	}
	g.Dump(user)
}
//自动识别转换Map数组
func scan_maps() {
	var (
		users  []map[string]string
		params = g.Slice{
			g.Map{
				"uid":  1,
				"name": "john",
			},
			g.Map{
				"uid":  2,
				"name": "smith",
			},
		}
	)

	if err := gconv.Scan(params, &users); err != nil {
		panic(err)
	}
	g.Dump(users)
}

