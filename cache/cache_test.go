package cache2go_test

import (
	"fmt"
	"lucky-golang/cache"
	"testing"
	"time"
)

// Keys & values in cache2go can be off arbitrary types, e.g. a struct.
type myStruct struct {
	text string

	moreData []byte
}

/***
go test -run='TestCache' -v
*/
/**
Golang单元测试对文件名和方法名，参数都有很严格的要求。
http://blog.csdn.net/samxx8/article/details/46894587
　　例如：
　　1、文件名必须以xx_test.go命名
　　2、方法必须是Test[^a-z]开头
　　3、方法参数必须 t *testing.T
　　之前就因为第 2 点没有写对，导致找了半天错误。现在真的让人记忆深刻啊，小小的东西当初看书没仔细。
　　下面分享一点go test的参数解读。来源
*/
func TestCache(t *testing.T) {
	// Accessing a new cache table for the first time will create it.
	cache := cache2go.Cache("myCache")

	// We will put a new item in the cache. It will expire after
	// not being accessed via Value(key) for more than 5 seconds.
	val := myStruct{"This is a test!", []byte{}}
	cache.Add("someKey", 5*time.Second, &val)

	// Let's retrieve the item from the cache.
	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Wait for the item to expire in cache.
	time.Sleep(6 * time.Second)
	res, err = cache.Value("someKey")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	// Add another item that never expires.
	cache.Add("someKey", 0, &val)

	// cache2go supports a few handy callbacks and loading mechanisms.
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache.
	cache.Delete("someKey")

	// And wipe the entire cache table.
	cache.Flush()
}
