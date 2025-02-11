package inmemory

import "github.com/joesweeny/midnite/internal/app"

type eventRepository struct {
	events []*app.Event
}

func (r *eventRepository) Insert(e *app.Event) error {
	r.events = append(r.events, e)
	return nil
}

func (r *eventRepository) ByUserID(id int) ([]*app.Event, error) {
	events := make([]*app.Event, 0)

	for _, event := range r.events {
		if event.UserID == id {
			events = append(events, event)
		}
	}

	return events, nil
}

func NewEventRepository() app.EventRepository {
	return &eventRepository{events: make([]*app.Event, 0)}
}
