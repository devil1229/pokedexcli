package pokecache

import (
	//"fmt"
	"time"
)

//method to add data to cache
func (sm *Cache) Add(key string, value []byte) {
	//fmt.Println("inside add")
	sm.mu.Lock()
	defer sm.mu.Unlock()
    //fmt.Println("Used Cache Add")
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	sm.cache[key] = entry
}

//method to retrive data from cache
func (sm *Cache) Get(key string) ([]byte, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
   // fmt.Println("Used Cache Get")
	value, exists := sm.cache[key]
	return value.val, exists
}


//just clean the cache on regular intervals so that we should not run into memory limit or use old api data
func (sm *Cache) realLoop() {
	
	sm.mu.Lock()
	//fmt.Println("Used Cache eviction")
	for key, val := range sm.cache {
		durationSinceCreation := time.Since(val.createdAt)
		if durationSinceCreation > sm.td {
			delete(sm.cache, key)
		}
	}
	sm.mu.Unlock()
	//time.Sleep(5 * time.Minute)
	
}