package memory

import (
	"calendar/internal/repository"
	"fmt"
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
			m.events[i] = event
		}
	}
	return nil
}

func (m *MemoryDb) DeleteEvent(id repository.Id) error {
	var newEvents []repository.Event

	for _, e := range m.events {
		if e.Id == id {
			continue
		} else {
			newEvents = append(newEvents, e)
		}
	}

	m.events = newEvents

	return nil
}

func filterDates(events []repository.Event, from time.Time, to time.Time) []repository.Event {
	var dayEvents []repository.Event

	for _, e := range events {
		fmt.Println("loop")
		fmt.Println(e.StartAt.After(from))
		fmt.Println(from)
		//fmt.Println(e.StartAt.Before(to))
		if (e.StartAt.After(from) || e.StartAt.Equal(from)) && e.StartAt.Before(to) {
			dayEvents = append(dayEvents, e)
		}
	}

	return dayEvents
}

func (m *MemoryDb) GetEventsDay(from time.Time) ([]repository.Event, error) {
	return filterDates(m.events, from, from.AddDate(0, 0, 1)), nil
}

func (m *MemoryDb) GetEventsWeek(from time.Time) ([]repository.Event, error) {
	return filterDates(m.events, from, from.AddDate(0, 0, 7)), nil
}

func (m *MemoryDb) GetEventsMonth(from time.Time) ([]repository.Event, error) {
	return filterDates(m.events, from, from.AddDate(0, 1, 0)), nil
}

var _ repository.BaseRepo = (*MemoryDb)(nil)
