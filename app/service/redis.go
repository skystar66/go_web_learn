package service

import "github.com/gogf/gf/frame/g"

var RedisService = redisModel{}

type redisModel struct {
}

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
