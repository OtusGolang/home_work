package repository

import "context"

type BaseRepo interface {
	Connect(ctx context.Context) error
	Close() error
	GetBooks(ctx context.Context) ([]Event, error)
}

type Event struct {
	id string
}