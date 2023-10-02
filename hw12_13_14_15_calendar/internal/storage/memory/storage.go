package memorystorage

import "sync"

type Storage struct {
	// TODO
	mu sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{}
}

// TODO
