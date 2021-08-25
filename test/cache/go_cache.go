package main

import (
	"fmt"
	"github.com/gogf/gf/os/gcache"
	"time"
)

func main() {

	test_lru_cache()
}

func test_cache() {

	// 创建一个缓存对象，
	// 当然也可以便捷地直接使用gcache包方法
	c := gcache.New()

	// 设置缓存，不过期
	c.Set("k1", "v1", 0)

	// 获取缓存
	v, _ := c.Get("k1")
	fmt.Println(v)

	// 获取缓存大小
	n, _ := c.Size()
	fmt.Println(n)

	// 缓存中是否存在指定键名
	b, _ := c.Contains("k1")
	fmt.Println(b)

	// 删除并返回被删除的键值
	fmt.Println(c.Remove("k1"))

	// 关闭缓存对象，让GC回收资源
	c.Close()

}


func control_cache() {
	// 当键名不存在时写入，设置过期时间1000毫秒
	gcache.SetIfNotExist("k1", "v1", 1000*time.Millisecond)

	// 打印当前的键名列表
	keys, _ := gcache.Keys()
	fmt.Println(keys)

	// 打印当前的键值列表
	values, _ := gcache.Values()
	fmt.Println(values)

	// 获取指定键值，如果不存在时写入，并返回键值
	v, _ := gcache.GetOrSet("k2", "v2", 0)
	fmt.Println(v)

	// 打印当前的键值对
	data1, _ := gcache.Data()
	fmt.Println(data1)

	// 等待1秒，以便k1:v1自动过期
	time.Sleep(time.Second)

	// 再次打印当前的键值对，发现k1:v1已经过期，只剩下k2:v2
	data2, _ := gcache.Data()
	fmt.Println(data2)
}

func test_lru_cache() {

	// 设置LRU淘汰数量
	cache:=gcache.New(2)

	// 添加10个元素，不过期
	for i := 0; i < 10; i++ {
		cache.Set(i,i,0)
	}
	n, _ := cache.Size()
	fmt.Println(n)
	keys, _ := cache.Keys()
	fmt.Println(keys)

	// 读取键名1，保证该键名是优先保留
	v, _ := cache.Get(1)
	fmt.Println(v)

	// 等待一定时间后(默认1秒检查一次)，
	// 元素会被按照从旧到新的顺序进行淘汰
	time.Sleep(2*time.Second)
	n, _ = cache.Size()
	fmt.Println(n)
	keys, _ = cache.Keys()
	fmt.Println(keys)
}









