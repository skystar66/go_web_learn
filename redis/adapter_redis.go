package redis

import (
	"context"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/os/gcache"
	"time"
)

type Redis struct {
	redis *gredis.Redis
}

//newAdapterMemory 创建并返回一个新的内存缓存对象。
func NewRedis(redis *gredis.Redis) gcache.Adapter {
	return &Redis{
		redis: redis,
	}
}

// 使用<key>-<value>对设置sets缓存，<duration>后过期。
// 如果 <duration> == 0，它不会过期。
// 如果 <duration> < 0 或给定的 <value> 为 nil，则删除 <key>。
func (c *Redis) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	var err error
	if value == nil || duration < 0 {
		_, err = c.redis.Ctx(ctx).DoVar("DEL", key)
	} else {
		if duration == 0 {
			_, err = c.redis.Ctx(ctx).DoVar("SET", key, value)
		} else {
			_, err = c.redis.Ctx(ctx).DoVar("SETEX", key, uint64(duration.Seconds()), value)
		}
	}
	return err
}

// Update 更新 <key> 的值而不更改其过期时间并返回旧值。
// 如果 <key> 在缓存中不存在，则返回值 <exist> 为 false。
// 如果给定的 <value> 为零，则删除 <key>。
// 如果 <key> 在缓存中不存在，它什么都不做。
func (c *Redis) Update(ctx context.Context, key interface{}, value interface{}) (oldValue interface{}, exist bool, err error) {
	var (
		v           *gvar.Var
		oldDuration time.Duration
	)
	// TTL.
	v, err = c.redis.Ctx(ctx).DoVar("TTL", key)
	if err != nil {
		return
	}
	oldDuration = v.Duration()
	if oldDuration == -2 {
		// It does not exist.
		return
	}
	// Check existence.
	v, err = c.redis.Ctx(ctx).DoVar("GET", key)
	if err != nil {
		return
	}
	oldValue = v.Val()
	// DEL.
	if value == nil {
		_, err = c.redis.Ctx(ctx).DoVar("DEL", key)
		if err != nil {
			return
		}
		return
	}
	// Update the value.
	if oldDuration == -1 {
		_, err = c.redis.Ctx(ctx).DoVar("SET", key, value)
	} else {
		oldDuration *= time.Second
		_, err = c.redis.Ctx(ctx).DoVar("SETEX", key, uint64(oldDuration.Seconds()), value)
	}
	return oldValue, true, err
}
// UpdateExpire 更新 <key> 的过期时间并返回旧的过期时间值。
//
// 如果 <key> 在缓存中不存在，则返回 -1。
// 如果 <duration> < 0，则删除 <key>。
func (c *Redis) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	var (
		v *gvar.Var
	)
	// TTL.
	v, err = c.redis.Ctx(ctx).DoVar("TTL", key)
	if err != nil {
		return
	}
	oldDuration = v.Duration()
	if oldDuration == -2 {
		// It does not exist.
		oldDuration = -1
		return
	}
	oldDuration *= time.Second
	// DEL.
	if duration < 0 {
		_, err = c.redis.Ctx(ctx).Do("DEL", key)
		return
	}
	// Update the expire.
	if duration > 0 {
		_, err = c.redis.Ctx(ctx).Do("EXPIRE", key, uint64(duration.Seconds()))
	}
	// No expire.
	if duration == 0 {
		v, err = c.redis.Ctx(ctx).DoVar("GET", key)
		if err != nil {
			return
		}
		_, err = c.redis.Ctx(ctx).Do("SET", key, v.Val())
	}
	return
}

// GetExpire 在缓存中检索并返回 <key> 的过期时间。
//
// 如果 <key> 没有过期，则返回 0。
// 如果 <key> 在缓存中不存在，则返回 -1。
func (c *Redis) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	v, err := c.redis.Ctx(ctx).DoVar("TTL", key)
	if err != nil {
		return 0, err
	}
	switch v.Int() {
	case -1:
		return 0, nil
	case -2:
		return -1, nil
	default:
		return v.Duration() * time.Second, nil
	}
}
// SetIfNotExist 使用 <key>-<value> 对设置缓存，该对在 <duration> 后过期
// 如果 <key> 在缓存中不存在。 它返回 true <key> 不存在于
// 缓存并将 <value> 成功设置到缓存中，否则返回 false。
//
// 参数 <value> 可以是 <func() interface{}> 的类型，但如果它是
// 结果为零。
//
// 如果 <duration> == 0，它不会过期。
// 如果 <duration> < 0 或给定的 <value> 为 nil，则删除 <key>。
func (c *Redis) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	var err error
	// Execute the function and retrieve the result.
	if f, ok := value.(func() (interface{}, error)); ok {
		value, err = f()
		if value == nil {
			return false, err
		}
	}
	// DEL.
	if duration < 0 || value == nil {
		v, err := c.redis.Ctx(ctx).DoVar("DEL", key, value)
		if err != nil {
			return false, err
		}
		if v.Int() == 1 {
			return true, err
		} else {
			return false, err
		}
	}
	v, err := c.redis.Ctx(ctx).DoVar("SETNX", key, value)
	if err != nil {
		return false, err
	}
	if v.Int() > 0 && duration > 0 {
		// Set the expire.
		_, err := c.redis.Ctx(ctx).Do("EXPIRE", key, uint64(duration.Seconds()))
		if err != nil {
			return false, err
		}
		return true, err
	}
	return false, err
}

// 使用<data>的键值对设置批处理集缓存，在<duration>后过期。
//
// 如果 <duration> == 0，它不会过期。
// 如果 <duration> < 0 或给定的 <value> 为 nil，则删除 <data> 的键。
func (c *Redis) Sets(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	if len(data) == 0 {
		return nil
	}
	// DEL.
	if duration < 0 {
		var (
			index = 0
			keys  = make([]interface{}, len(data))
		)
		for k, _ := range data {
			keys[index] = k
			index += 1
		}
		_, err := c.redis.Ctx(ctx).Do("DEL", keys...)
		if err != nil {
			return err
		}
	}
	if duration == 0 {
		var (
			index     = 0
			keyValues = make([]interface{}, len(data)*2)
		)
		for k, v := range data {
			keyValues[index] = k
			keyValues[index+1] = v
			index += 2
		}
		_, err := c.redis.Ctx(ctx).Do("MSET", keyValues...)
		if err != nil {
			return err
		}
	}
	if duration > 0 {
		var err error
		for k, v := range data {
			if err = c.Set(ctx, k, v, duration); err != nil {
				return err
			}
		}
	}
	return nil
}

// Get 检索并返回给定 <key> 的关联值。
// 如果它不存在或者它的值为nil，则返回nil。
func (c *Redis) Get(ctx context.Context, key interface{}) (interface{}, error) {
	v, err := c.redis.Ctx(ctx).DoVar("GET", key)
	if err != nil {
		return nil, err
	}
	return v.Val(), nil
}

// GetOrSet retrieves and returns the value of <key>, or sets <key>-<value> pair and
// returns <value> if <key> does not exist in the cache. The key-value pair expires
// after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil, but it does nothing
// if <value> is a function and the function result is nil.
func (c *Redis) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (interface{}, error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return value, c.Set(ctx, key, value, duration)
	} else {
		return v, nil
	}
}

// GetOrSetFunc retrieves and returns the value of <key>, or sets <key> with result of
// function <f> and returns its result if <key> does not exist in the cache. The key-value
// pair expires after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil, but it does nothing
// if <value> is a function and the function result is nil.
func (c *Redis) GetOrSetFunc(ctx context.Context, key interface{}, f func() (interface{}, error), duration time.Duration) (interface{}, error) {
	v, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		value, err := f()
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return value, c.Set(ctx, key, value, duration)
	} else {
		return v, nil
	}
}

// GetOrSetFuncLock retrieves and returns the value of <key>, or sets <key> with result of
// function <f> and returns its result if <key> does not exist in the cache. The key-value
// pair expires after <duration>.
//
// It does not expire if <duration> == 0.
// It does nothing if function <f> returns nil.
//
// Note that the function <f> should be executed within writing mutex lock for concurrent
// safety purpose.
func (c *Redis) GetOrSetFuncLock(ctx context.Context, key interface{}, f func() (interface{}, error), duration time.Duration) (interface{}, error) {
	return c.GetOrSetFunc(ctx, key, f, duration)
}

// Contains returns true if <key> exists in the cache, or else returns false.
func (c *Redis) Contains(ctx context.Context, key interface{}) (bool, error) {
	v, err := c.redis.Ctx(ctx).DoVar("EXISTS", key)
	if err != nil {
		return false, err
	}
	return v.Bool(), nil
}

// Remove deletes the one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the deleted last item.
func (c *Redis) Remove(ctx context.Context, keys ...interface{}) (value interface{}, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	// Retrieves the last key value.
	if v, err := c.redis.Ctx(ctx).DoVar("GET", keys[len(keys)-1]); err != nil {
		return nil, err
	} else {
		value = v.Val()
	}
	// Deletes all given keys.
	_, err = c.redis.Ctx(ctx).DoVar("DEL", keys...)
	return value, err
}

// Data returns a copy of all key-value pairs in the cache as map type.
func (c *Redis) Data(ctx context.Context) (map[interface{}]interface{}, error) {
	// Keys.
	v, err := c.redis.Ctx(ctx).DoVar("KEYS", "*")
	if err != nil {
		return nil, err
	}
	keys := v.Slice()
	// Values.
	v, err = c.redis.Ctx(ctx).DoVar("MGET", keys...)
	if err != nil {
		return nil, err
	}
	values := v.Slice()
	// Compose keys and values.
	data := make(map[interface{}]interface{})
	for i := 0; i < len(keys); i++ {
		data[keys[i]] = values[i]
	}
	return data, nil
}

// Keys returns all keys in the cache as slice.
func (c *Redis) Keys(ctx context.Context) ([]interface{}, error) {
	v, err := c.redis.Ctx(ctx).DoVar("KEYS", "*")
	if err != nil {
		return nil, err
	}
	return v.Slice(), nil
}

// Values returns all values in the cache as slice.
func (c *Redis) Values(ctx context.Context) ([]interface{}, error) {
	// Keys.
	v, err := c.redis.Ctx(ctx).DoVar("KEYS", "*")
	if err != nil {
		return nil, err
	}
	keys := v.Slice()
	// Values.
	v, err = c.redis.Ctx(ctx).DoVar("MGET", keys...)
	if err != nil {
		return nil, err
	}
	return v.Slice(), nil
}

// Size returns the size of the cache.
func (c *Redis) Size(ctx context.Context) (size int, err error) {
	v, err := c.redis.Ctx(ctx).DoVar("DBSIZE")
	if err != nil {
		return 0, err
	}
	return v.Int(), nil
}

// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
func (c *Redis) Clear(ctx context.Context) error {
	// The "FLUSHDB" may not be available.
	if _, err := c.redis.Ctx(ctx).DoVar("FLUSHDB"); err != nil {
		keys, err := c.Keys(ctx)
		if err != nil {
			return err
		}
		_, err = c.Remove(ctx, keys...)
		return err
	}
	return nil
}

// Close closes the cache.
func (c *Redis) Close(ctx context.Context) error {
	// It does nothing.
	return nil
}











