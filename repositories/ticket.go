package repositories

import (
	"context"

	"github.com/joaolucassilvadev/Ticket-booking-app/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Preload("Event").First(&ticket)

	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil // Retorna nil quando não encontrado
		}
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(ticket).Create(&ticket)
	if res.Error != nil {
		return nil, res.Error
	}
	return r.GetOne(ctx, ticket.ID)
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updatedata map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}
	updateRes := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updatedata)
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}
	return r.GetOne(ctx, ticket.ID)
}
