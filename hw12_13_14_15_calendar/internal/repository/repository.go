package repository

import (
	"time"
)

type BaseRepo interface {
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
	StartAt     time.Time
	EndAt       time.Time
	Description string
	UserId      Id
	NotifyAt    time.Time
}

type User struct {
	Id
}
