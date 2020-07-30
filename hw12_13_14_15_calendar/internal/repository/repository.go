package repository

import (
	"context"
	"time"
)

type BaseRepo interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	GetEvents(ctx context.Context) ([]Event, error)
	AddEvent(event Event) error
	UpdateEvent(event Event) error
	DeleteEvent(id Id) error
	GetEventsDay(from time.Time) ([]Event, error)
	GetEventsWeek(from time.Time) ([]Event, error)
	GetEventsMonth(from time.Time) ([]Event, error)
}

type Id = int

type Event struct {
	Id
	Title       string
	DateTime    time.Time
	StartAt     time.Time
	EndAt       time.Time
	Description string
	UserId      Id
	NotifyAt    time.Time
}

type User struct {
	Id
}
