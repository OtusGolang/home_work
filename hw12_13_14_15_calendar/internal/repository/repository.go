package repository

import (
	"context"
	"time"
)

type BaseRepo interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	AddEvent(event Event) error
	UpdateEvent(event Event) error
	DeleteEvent(userId Id, eventId Id) error
	GetEventsDay(userId Id, from time.Time) ([]Event, error)
	GetEventsWeek(userId Id, from time.Time) ([]Event, error)
	GetEventsMonth(userId Id, from time.Time) ([]Event, error)
}

type Id = int

type Event struct {
	Id          Id
	Title       string
	StartAt     time.Time `db:"start_at"`
	EndAt       time.Time `db:"end_at"`
	Description string
	UserId      int       `db:"user_id"`
	NotifyAt    time.Time `db:"notify_at"`
}

type User struct {
	Id
}
