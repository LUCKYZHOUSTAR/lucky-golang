package cache2go

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex
	key      interface{}
	data     interface{}
	lifeSpan time.Duration

	createdOn time.Time

	accessdOn     time.Time
	accessCount   int64
	aboutToExpire func(key interface{})
}

//大写的开头，在包外可以访问的到
func NewCacheItem(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key:           key,
		lifeSpan:      lifeSpan,
		createdOn:     t,
		accessdOn:     t,
		accessCount:   0,
		aboutToExpire: nil,
		data:          data,
	}
}

func (item *CacheItem) keepAlive() {
	item.Lock()
	defer item.Unlock()
	item.accessdOn = time.Now()
	item.accessCount++
}

func (item *CacheItem) LifeSpan() time.Duration {
	return item.lifeSpan
}
func (item *CacheItem) AccessedOn() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.accessdOn
}
func (item *CacheItem) CreatedOn() time.Time {
	// immutable
	return item.createdOn
}
func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}
func (item *CacheItem) Key() interface{} {
	// immutable
	return item.key
}
func (item *CacheItem) Data() interface{} {
	// immutable
	return item.data
}
func (item *CacheItem) SetAboutToExpireCallback(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = f
}
