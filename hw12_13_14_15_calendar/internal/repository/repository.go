package repository

import (
	"time"
)

type BaseRepo interface {
	AddEvent(event Event) error
	UpdateEvent(event Event) error
	DeleteEvent(userId Id, eventId Id) error
	GetEventsDay(userId Id, from time.Time) ([]Event, error)
	GetEventsWeek(userId Id, from time.Time) ([]Event, error)
	GetEventsMonth(userId Id, from time.Time) ([]Event, error)
}

type Id = int

type Event struct {
	Id
	Title       string
	StartAt     time.Time
	EndAt       time.Time
	Description string
	UserId      Id
	NotifyAt    time.Time
}

type User struct {
	Id
}
