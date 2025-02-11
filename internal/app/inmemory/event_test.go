package inmemory_test

import (
	"github.com/joesweeny/midnite/internal/app/inmemory"
	"testing"

	"github.com/joesweeny/midnite/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	t.Run("inserts event records", func(t *testing.T) {
		t.Helper()

		repo := inmemory.NewEventRepository()

		event := &app.Event{
			Type:   "Bet",
			Amount: "100",
			UserID: 1,
			T:      12345,
		}

		err := repo.Insert(event)

		assert.NoError(t, err, "Expected no error when inserting an event")

		events, err := repo.ByUserID(1)

		assert.NoError(t, err, "Expected no error when fetching events by user ID")
		assert.Len(t, events, 1, "Expected exactly one event for user ID 1")
		assert.Equal(t, event, events[0], "The event returned does not match the inserted event")
	})
}

func TestByUserID(t *testing.T) {
	t.Run("returns events for a user id", func(t *testing.T) {
		t.Helper()

		repo := inmemory.NewEventRepository()

		event1 := &app.Event{
			Type:   "Bet",
			Amount: "50",
			UserID: 1,
			T:      12345,
		}
		event2 := &app.Event{
			Type:   "Bet",
			Amount: "200",
			UserID: 2,
			T:      12346,
		}
		event3 := &app.Event{
			Type:   "Bet",
			Amount: "150",
			UserID: 1,
			T:      12347,
		}

		// Insert the events
		_ = repo.Insert(event1)
		_ = repo.Insert(event2)
		_ = repo.Insert(event3)

		events, err := repo.ByUserID(1)
		assert.NoError(t, err, "Expected no error when fetching events by user ID")
		assert.Len(t, events, 2, "Expected two events for user ID 1")

		assert.Contains(t, events, event1, "Event1 should be returned for user ID 1")
		assert.Contains(t, events, event3, "Event3 should be returned for user ID 1")

		// Fetch events for user ID 2
		events, err = repo.ByUserID(2)
		assert.NoError(t, err, "Expected no error when fetching events by user ID")
		assert.Len(t, events, 1, "Expected one event for user ID 2")
		assert.Equal(t, event2, events[0], "The event returned does not match event2")
	})

	t.Run("returns no events if they do not exist for a user", func(t *testing.T) {
		t.Helper()

		repo := inmemory.NewEventRepository()

		events, err := repo.ByUserID(99)
		assert.NoError(t, err, "Expected no error when fetching events by user ID")
		assert.Len(t, events, 0, "Expected no events for non-existing user ID")
	})
}
