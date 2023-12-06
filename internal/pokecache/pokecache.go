package pokecache

import "time"

func NewCache(d time.Duration) *Cache {

	memoryCache := Cache{
		cache: map[string]cacheEntry{},
		td:    d,
	}
	// Create a ticker that ticks every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)

	// Start a goroutine to run the function every time the ticker ticks
	go func() {
		for {
			select {
			case <-ticker.C:
				memoryCache.realLoop()
			}
		}
	}()

	return &memoryCache

}