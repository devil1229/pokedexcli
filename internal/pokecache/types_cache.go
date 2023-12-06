package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
	td time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

