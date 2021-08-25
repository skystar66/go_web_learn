package main

import (
	"fmt"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"reflect"
)

type User struct {
	Id         int
	Passport   string
	Password   string
	Nickname   string
	CreateTime *gtime.Time
}

func (user *User) UnmarshalValue(value interface{}) error {
	if record, ok := value.(gdb.Record); ok {
		*user = User{
			Id:         record["id"].Int(),
			Passport:   record["passport"].String(),
			Password:   record["password"].String(),
			Nickname:   record["nickname"].String(),
			CreateTime: record["create_time"].GTime(),
		}
		return nil
	}
	return gerror.Newf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}

func main() {
	var (
		err   error
		users []*User
	)
	//创建数组
	array := garray.New(true)
	for i := 1; i <= 10; i++ {
		array.Append(g.Map{
			"id":       i,
			"passport": fmt.Sprintf(`user_%d`, i),
			"password": fmt.Sprintf(`pass_%d`, i),
			"nickname": fmt.Sprintf(`name_%d`, i),
		})
	}
	//写入数据
	if _, err := g.DB().Model("userconv").Data(array).Insert(); err != nil {
		panic(err)
	}
	//查询数据
	err = g.DB().Model("userconv").Order("id asc").Scan(&users)
	if err != nil {
		panic(err)
	}
	g.Dump(users)
}
