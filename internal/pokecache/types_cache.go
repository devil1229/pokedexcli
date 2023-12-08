package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex  // adding the mutex lock is important so that you don't run into race condition
	td time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

