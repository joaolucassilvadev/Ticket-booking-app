package repositories

import (
	"context"
	"time"

	"github.com/joaolucassilvadev/Ticket-booking-app/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	evensts := []*models.Event{}

	evensts = append(evensts, &models.Event{
		Id:        "87876287628721",
		Name:      "show",
		Location:  "casa",
		Date:      time.Now(),
		CreatedAt: time.Now(),
	})

	return evensts, nil
}
func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}
func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}
func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
