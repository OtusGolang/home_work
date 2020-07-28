package repository

import "context"

type BaseRepo interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	GetEvents(ctx context.Context) ([]Event, error)
}

type Event struct {
	Id int
	Title int
	Body int
}