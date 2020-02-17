package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	// Place your code here
}

type lruCache struct {
	// Place your code here:
	// - capacity
	// - queue
	// - items
}

type cacheItem struct {
	// Place your code here
}

func NewCache(capacity int) Cache {
	return &lruCache{}
}
