package memory

import (
	"calendar/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var userId = 1
var wrongUserId = 2

func TestImMemoryImplementation(t *testing.T) {
	t.Run("test add", func(t *testing.T) {
		db := new(MemoryDb)
		startTime := time.Now()

		dbEvents, _ := db.GetEventsDay(userId, startTime)
		assert.Equal(t, 0, len(dbEvents))

		_ = db.AddEvent(createEvent(startTime))
		dbEvents, _ = db.GetEventsDay(userId, startTime)
		assert.Equal(t, 1, len(dbEvents))
	})

	t.Run("test update", func(t *testing.T) {
		db := new(MemoryDb)
		startTime := time.Now()
		initialEvent := createEvent(startTime)

		_ = db.AddEvent(initialEvent)
		updatedEvent := initialEvent
		updatedEvent.Title = "updated"
		err := db.UpdateEvent(updatedEvent)

		dbEvents, _ := db.GetEventsDay(userId, startTime)

		assert.NoError(t, err)
		assert.Equal(t, initialEvent.Id, dbEvents[0].Id)
		assert.Equal(t, updatedEvent.Title, dbEvents[0].Title)
	})

	t.Run("test update, auth error", func(t *testing.T) {
		db := new(MemoryDb)
		startTime := time.Now()
		initialEvent := createEvent(startTime)

		_ = db.AddEvent(initialEvent)
		updatedEvent := initialEvent
		updatedEvent.UserId = wrongUserId
		updatedEvent.Title = "updated"
		err := db.UpdateEvent(updatedEvent)

		assert.Error(t, err)
	})

	t.Run("test delete", func(t *testing.T) {
		db := new(MemoryDb)
		initialEvent := createEvent(time.Now())

		_ = db.AddEvent(initialEvent)

		err := db.DeleteEvent(userId, initialEvent.Id)

		dbEvents, _ := db.GetEventsDay(userId, time.Now())

		assert.NoError(t, err)
		assert.Equal(t, 0, len(dbEvents))
	})

	t.Run("test delete, auth error", func(t *testing.T) {
		db := new(MemoryDb)
		initialEvent := createEvent(time.Now())

		_ = db.AddEvent(initialEvent)

		err := db.DeleteEvent(wrongUserId, initialEvent.Id)

		assert.Error(t, err)
	})

	t.Run("test get events, day", func(t *testing.T) {
		db := new(MemoryDb)

		threshold := time.Now()
		for _, d := range []time.Duration{-3, -2, -1, 0, 1, 2, 3, 25, 26, 27} {
			event := createEvent(threshold.Add(time.Hour * d))
			_ = db.AddEvent(event)
		}

		dbEvents, _ := db.GetEventsDay(userId, threshold)
		assert.Equal(t, 4, len(dbEvents))
	})

	t.Run("test get events, week", func(t *testing.T) {
		db := new(MemoryDb)

		threshold := time.Now()
		for _, d := range []time.Duration{-3, -2, -1, 0, 1, 2, 3, 25, 26, 27} {
			event := createEvent(threshold.Add(time.Hour * 24 * d))
			_ = db.AddEvent(event)
		}

		dbEvents, _ := db.GetEventsWeek(userId, threshold)
		assert.Equal(t, 4, len(dbEvents))
	})

	t.Run("test get events, month", func(t *testing.T) {
		db := new(MemoryDb)

		threshold := time.Now()
		for _, d := range []time.Duration{-3, -2, -1, 0, 1, 2, 3, 25, 26, 27} {
			week := time.Hour * 24 * 7
			event := createEvent(threshold.Add(week * d))
			_ = db.AddEvent(event)
		}

		dbEvents, _ := db.GetEventsMonth(userId, threshold)
		assert.Equal(t, 4, len(dbEvents))
	})
}

func createEvent(initialTime time.Time) repository.Event {
	return repository.Event{
		Id:          0,
		Title:       "title",
		StartAt:     initialTime,
		EndAt:       initialTime.Add(time.Hour * 24),
		Description: "description",
		UserId:      userId,
		NotifyAt:    initialTime.Add(time.Hour * -24),
	}
}
