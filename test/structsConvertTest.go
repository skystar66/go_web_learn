package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {

	type User struct {
		Uid  int
		Name string
	}
	params:=g.Slice{
		g.Map{
			"uid":1,
			"name":"xl",
		},
		g.Map{
			"uid":2,
			"name":"zy",
		},
	}

	var users []*User
	if err:=gconv.Structs(params,&users);err!=nil {
		fmt.Println(err.Error())
	}else {
		g.Dump(users)
	}
}

