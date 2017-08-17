package cache2go

import (
	"log"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex
	name            string
	items           map[interface{}]*CacheItem
	cleanUpTimer    *time.Timer
	cleanUpInterval time.Duration
	logger          *log.Logger

	loadData          func(key interface{}, args ...interface{}) *CacheItem
	addedItem         func(item *CacheItem)
	aboutToDeleteItem func(item *CacheItem)
}

func (table *CacheTable) Count() int {
	table.Lock()
	defer table.Unlock()
	return len(table.items)
}

func (table *CacheTable) Foreach(trans func(key interface{}, item *CacheItem)) {
	table.RLock()
	defer table.RUnlock()
	for k, v := range table.items {
		trans(k, v)
	}
}

func (table *CacheTable) SetDataLoader(f func(interface{}, ...interface{}) *CacheItem) {
	table.Lock()
	defer table.Unlock()
	table.loadData = f
}

func (table *CacheTable) SetAddedItemCallBack(f func(item *CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.addedItem = f
}

// SetAboutToDeleteItemCallback configures a callback, which will be called
// every time an item is about to be removed from the cache.
func (table *CacheTable) SetAboutToDeleteItemCallback(f func(*CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.aboutToDeleteItem = f
}

// SetLogger sets the logger to be used by this cache table.
func (table *CacheTable) SetLogger(logger *log.Logger) {
	table.Lock()
	defer table.Unlock()
	table.logger = logger
}

// Expiration check loop, triggered by a self-adjusting timer.
func (table *CacheTable) expirationCheck() {
	table.Lock()
	if table.cleanUpTimer != nil {
		table.cleanUpTimer.Stop()
	}

	if table.cleanUpInterval > 0 {
		table.log("Expiration check triggered after", table.cleanUpInterval, "for table", table.name)

	} else {
		table.log("Expiration check installed for table", table.name)
	}

	// Cache value so we don't keep blocking the mutex.
	items := table.items
	table.Unlock()
	// To be more accurate with timers, we would need to update 'now' on every
	// loop iteration. Not sure it's really efficient though.
	smallestDuration := 0 * time.Second
	for key, item := range items {
		// Cache values so we don't keep blocking the mutex.
		item.RLock()
		lifeSpan := item.lifeSpan
		accessedOn := item.accessdOn
		item.RUnlock()
		if lifeSpan == 0 {
			continue
		}

		now := time.Now()
		if now.Sub(accessedOn) >= lifeSpan {
			table.Delete(key)
		} else {
			if smallestDuration == 0 || lifeSpan-now.Sub(accessedOn) < smallestDuration {
				smallestDuration = lifeSpan - now.Sub(accessedOn)
			}
		}

	}

	table.Lock()
	table.cleanUpInterval = smallestDuration
	if smallestDuration > 0 {
		table.cleanUpTimer = time.AfterFunc(smallestDuration, func() {
			go table.expirationCheck()
		})
	}
	table.Unlock()
}

func (table *CacheTable) Delete(key interface{}) (*CacheItem, error) {
	table.Lock()
	r, ok := table.items[key]
	if !ok {
		table.Unlock()
		return nil, ErrKeyNotFound
	}
	// Cache value so we don't keep blocking the mutex.
	aboutToDeleteItem := table.aboutToDeleteItem
	table.Unlock()
	if aboutToDeleteItem != nil {
		// Trigger callbacks before deleting an item from cache.
		aboutToDeleteItem(r)
	}

	r.RLock()
	defer r.RUnlock()
	if r.aboutToExpire != nil {
		r.aboutToExpire(key)
	}
	table.log("Deleting item with key", key, "created on", r.createdOn, "and hit", r.accessCount, "times from table", table.name)
	delete(table.items, key)

	return r, nil
}

func (table *CacheTable) Exists(key interface{}) bool {
	table.RLock()
	defer table.Unlock()
	_, ok := table.items[key]
	return ok
}

func (table *CacheTable) addInternal(item *CacheItem) {
	table.log("Adding item with key", item.key, "and lifespan of", item.lifeSpan, "to table", table.name)
	//add key to items
	table.items[item.key] = item
	// Cache values so we don't keep blocking the mutex.
	expDur := table.cleanUpInterval
	addedItem := table.addedItem
	table.Unlock()
	// Trigger callback after adding an item to cache.
	if addedItem != nil {
		addedItem(item)
	}

	// If we haven't set up any expiration check timer or found a more imminent item.
	//lifespan==0 show the key never expiration
	if item.lifeSpan > 0 && (expDur == 0 || item.lifeSpan < expDur) {
		table.expirationCheck()
	}
}
func (table *CacheTable) Add(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	item := NewCacheItem(key, lifeSpan, data)

	// Add item to cache.
	table.Lock()
	table.addInternal(item)

	return item
}

func (table *CacheTable) NotFoundAdd(key interface{}, lifeSpan time.Duration, data interface{}) bool {
	table.Lock()

	if _, ok := table.items[key]; ok {
		table.Unlock()
		return false
	}

	item := NewCacheItem(key, lifeSpan, data)
	table.addInternal(item)

	return true
}

func (table *CacheTable) Value(key interface{}, args ...interface{}) (*CacheItem, error) {
	table.RLock()
	r, ok := table.items[key]
	loadData := table.loadData
	table.RUnlock()

	if ok {
		// Update access counter and timestamp.
		r.keepAlive()
		return r, nil
	}

	// Item doesn't exist in cache. Try and fetch it with a data-loader.
	if loadData != nil {
		item := loadData(key, args...)
		if item != nil {
			table.Add(key, item.lifeSpan, item.data)
			return item, nil
		}

		return nil, ErrKeyNotFoundOrLoadable
	}

	return nil, ErrKeyNotFound
}

// Flush deletes all items from this cache table.
func (table *CacheTable) Flush() {
	table.Lock()
	defer table.Unlock()

	table.log("Flushing table", table.name)

	//clean all items
	table.items = make(map[interface{}]*CacheItem)
	table.cleanUpInterval = 0
	if table.cleanUpTimer != nil {
		table.cleanUpTimer.Stop()
	}
}
func (table *CacheTable) log(v ...interface{}) {
	if table.logger == nil {
		return
	}
	table.logger.Println(v)
}
