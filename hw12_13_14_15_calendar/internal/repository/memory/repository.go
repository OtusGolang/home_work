package memory

import (
	"calendar/internal/repository"
	"context"
	"errors"
	"time"
)

type DB struct {
	events []repository.Event
}

func (m *DB) Connect(ctx context.Context, dsn string) error {
	return nil
}

func (m *DB) Close() error {
	return nil
}

func (m *DB) AddEvent(event repository.Event) error {
	m.events = append(m.events, event)

	return nil
}

func (m *DB) UpdateEvent(event repository.Event) error {
	for i, e := range m.events {
		if e.ID == event.ID {
			if e.UserID != event.UserID {
				return errors.New("unauthorized request")
			}

			m.events[i] = event
		}
	}

	return nil
}

func (m *DB) DeleteEvent(userID repository.ID, eventID repository.ID) error {
	var newEvents []repository.Event

	for _, e := range m.events {
		if e.ID == eventID {
			if e.UserID != userID {
				return errors.New("unauthorized request")
			}

			continue
		} else {
			newEvents = append(newEvents, e)
		}
	}

	m.events = newEvents

	return nil
}

func filterDates(userID repository.ID, events []repository.Event, from time.Time, to time.Time) []repository.Event {
	var dayEvents []repository.Event

	for _, e := range events {
		if e.UserID == userID && (e.StartAt.After(from) || e.StartAt.Equal(from)) && e.StartAt.Before(to) {
			dayEvents = append(dayEvents, e)
		}
	}

	return dayEvents
}

func (m *DB) GetEventsDay(userID repository.ID, from time.Time) ([]repository.Event, error) {
	return filterDates(userID, m.events, from, from.AddDate(0, 0, 1)), nil
}

func (m *DB) GetEventsWeek(userID repository.ID, from time.Time) ([]repository.Event, error) {
	return filterDates(userID, m.events, from, from.AddDate(0, 0, 7)), nil
}

func (m *DB) GetEventsMonth(userID repository.ID, from time.Time) ([]repository.Event, error) {
	return filterDates(userID, m.events, from, from.AddDate(0, 1, 0)), nil
}
