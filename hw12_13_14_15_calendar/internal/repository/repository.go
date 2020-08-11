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
	DeleteEvent(userID ID, eventId ID) error
	GetEventsDay(userID ID, from time.Time) ([]Event, error)
	GetEventsWeek(userID ID, from time.Time) ([]Event, error)
	GetEventsMonth(userID ID, from time.Time) ([]Event, error)
}

type ID = int

type Event struct {
	Id          ID
	Title       string
	StartAt     time.Time `db:"start_at"`
	EndAt       time.Time `db:"end_at"`
	Description string
	UserID      int       `db:"user_id"`
	NotifyAt    time.Time `db:"notify_at"`
}

type User struct {
	ID
}
