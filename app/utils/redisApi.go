package utils

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
)

var RediApi = redisApi{}

type redisApi struct {
}

var (
	config = gredis.Config{
		Host: "127.0.0.1",
		Port: 6379,
		Db:   0,
	}
	group = "test"
	redis = &gredis.Redis{}
)

func init() {
	gredis.SetConfig(&config,group)
	redis =  gredis.Instance(group)
	g.Log().Infof("High Redis Init Success! group : %s", group)

}

func (receiver *redisApi) Get(key string) string  {
	val,_:=redis.DoVar("GET",key)
	g.Log().Infof("High Redis Get key: %s value: %s", key, val.String())
	return val.String()
}


func (receiver *redisApi) Set(key string,value string) error  {
	_,err:=redis.DoVar("SET",key,value)
	g.Log().Infof("High Redis Set key: %s value: %s", key, value)
	return err
}



