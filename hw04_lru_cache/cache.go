package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	// Place your code here
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]ListItem, capacity),
	}
}
