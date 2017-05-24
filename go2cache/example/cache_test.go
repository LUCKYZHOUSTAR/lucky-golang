package go2cache__test

import (
	"fmt"
	"lucky-golang/go2cache"
	"strconv"
	"testing"
	"time"
)

// Keys & values in go2cache can be off arbitrary types, e.g. a struct.
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

	cache := go2cache.Cache("myCache")

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

	// go2cache supports a few handy callbacks and loading mechanisms.
	cache.SetAboutToDeleteItemCallback(func(e *go2cache.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache.
	cache.Delete("someKey")

	// And wipe the entire cache table.
	cache.Flush()
}

// test to load data from the cache
func TestDataLoader(t *testing.T) {

	cache := go2cache.Cache("myCache")

	//if the key not exists in the cache ,so will fetch the key  form the loaddata
	cache.SetDataLoader(func(key interface{}, args ...interface{}) *go2cache.CacheItem {

		val := "this is a test with key" + key.(string)
		item := go2cache.NewCacheItem(key, 0, val)

		return item
	})

	for i := 0; i < 10; i++ {
		res, err := cache.Value("somekey_" + strconv.Itoa(i))
		if err == nil {
			fmt.Println("Found value in cazhe", res.Data())
		} else {
			fmt.Println("Error retrieving value from cache:", err)
		}
	}
}

func TestCallback(t *testing.T) {

	cache := go2cache.Cache("mycache")

	cache.SetAddedItemCallback(func(entry *go2cache.CacheItem) {
		fmt.Println("Added :", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.SetAboutToDeleteItemCallback(func(entry *go2cache.CacheItem) {
		fmt.Println("Deleting:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.Add("some key ", 0, "this is a test")

	res, err := cache.Value("some key")
	if err == nil {
		fmt.Println("Found value in the cache", res.Data())
	}

	fmt.Println("Error retrieving value from cache:", err)

	cache.Delete("some key")
	// Caching a new item that expires in 3 seconds
	res = cache.Add("anotherKey", 3*time.Second, "This is another test")

	// This callback will be triggered when the item is about to expire
	res.SetAboutToExpireCallback(func(key interface{}) {
		fmt.Println("About to expire:", key.(string))
	})

	time.Sleep(5 * time.Second)
}
