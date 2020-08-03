package memory

import (
	"calendar/internal/repository"
	"errors"
	"time"
)

type MemoryDb struct {
	events []repository.Event
}

func (m *MemoryDb) AddEvent(event repository.Event) error {
	m.events = append(m.events, event)
	return nil
}

func (m *MemoryDb) UpdateEvent(event repository.Event) error {
	for i, e := range m.events {
		if e.Id == event.Id {
			if e.UserId != event.UserId {
				return errors.New("Unauthorized request")
			}

			m.events[i] = event
		}
	}
	return nil
}

func (m *MemoryDb) DeleteEvent(userId repository.Id, eventId repository.Id) error {
	var newEvents []repository.Event

	for _, e := range m.events {
		if e.Id == eventId {
			if e.UserId != userId {
				return errors.New("Unauthorized request")
			}

			continue
		} else {
			newEvents = append(newEvents, e)
		}
	}

	m.events = newEvents

	return nil
}

func filterDates(userId repository.Id, events []repository.Event, from time.Time, to time.Time) []repository.Event {
	var dayEvents []repository.Event

	for _, e := range events {
		if e.UserId == userId && (e.StartAt.After(from) || e.StartAt.Equal(from)) && e.StartAt.Before(to) {
			dayEvents = append(dayEvents, e)
		}
	}

	return dayEvents
}

func (m *MemoryDb) GetEventsDay(userId repository.Id, from time.Time) ([]repository.Event, error) {
	return filterDates(userId, m.events, from, from.AddDate(0, 0, 1)), nil
}

func (m *MemoryDb) GetEventsWeek(userId repository.Id, from time.Time) ([]repository.Event, error) {
	return filterDates(userId, m.events, from, from.AddDate(0, 0, 7)), nil
}

func (m *MemoryDb) GetEventsMonth(userId repository.Id, from time.Time) ([]repository.Event, error) {
	return filterDates(userId, m.events, from, from.AddDate(0, 1, 0)), nil
}

var _ repository.BaseRepo = (*MemoryDb)(nil)
