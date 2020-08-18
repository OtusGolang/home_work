package memorystorage

import "sync"

type Storage struct {
	// TODO
	mu sync.RWMutex
}

func New() *Storage {
	return &Storage{}
}

// TODO
