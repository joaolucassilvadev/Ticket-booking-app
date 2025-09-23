package repositories

import (
	"context"

	"github.com/joaolucassilvadev/Ticket-booking-app/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := r.db.WithContext(ctx).Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	var event models.Event

	res := r.db.WithContext(ctx).Where("id = ?", eventId).First(&event)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil // Retorna nil quando não encontrado
		}
		return nil, res.Error
	}

	return &event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	res := r.db.WithContext(ctx).Create(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return &event, nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
