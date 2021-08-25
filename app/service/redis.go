package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

var RedisService = redisModel{}

type redisModel struct {
}

const (
	cache_map_key       = "cache_map_key"
	cache_mulit_map_key = "cache_mulit_map_key"
)

func (receiver *redisModel) GetVal(key string) string {
	value, _ := g.Redis().DoVar("GET", key)
	g.Log().Infof("Redis Get key: %s value: %s", key, value)
	return value.String()
}

func (receiver *redisModel) SetVal(key string, value string) error {
	_, err := g.Redis().DoVar("SET", key, value)
	g.Log().Infof("Redis Set key: %s value: %s", key, value)
	return err
}

func (receiver redisModel) Hset(filed string, value string) error {
	_, err := g.Redis().Do("HSET", cache_map_key, filed, value)
	g.Log().Infof("Redis hset key: %s value: %s", filed, value)

	if err != nil {
		panic(err)
		return err
	}
	return nil
}
func (receiver redisModel) HgetAll() g.Map {
	val, err := g.Redis().DoVar("HGETALL", cache_map_key)
	g.Log().Infof("Redis hget key: %s value: %s", cache_map_key, val.Map())
	if err != nil {
		panic(err)
		return nil
	}
	return val.Map()
}
func (receiver *redisModel) HMset(mapvals g.Map) error{
	_, err := g.Redis().Do("HMSET", append(g.Slice{cache_mulit_map_key}, gutil.MapToSlice(mapvals)...)...)
	g.Log().Infof("Redis hmset key: %s value: %s", cache_mulit_map_key, mapvals)

	if err != nil {
		panic(err)
		return err
	}
	return nil
}
func (receiver *redisModel) HMget(filed string) g.Slice{
	val,_:=g.Redis().DoVar("HMGET",cache_mulit_map_key,filed)
	g.Log().Infof("Redis hmget key: %s value: %s", filed, val.Slice())

	return val.Slice()
}
