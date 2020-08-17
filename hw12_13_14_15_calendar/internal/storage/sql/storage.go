package sqlstorage

import "context"

type Storage struct {
	// TODO
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
}
