package repository

import (
	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

func (repo *Repository) GetTicketList() ([]entity.AirTicket, error) {
	return nil, nil
}

func (repo *Repository) EditTicket(order uuid.UUID, edited entity.AirTicket) error {
	return nil
}

func (repo *Repository) RemoveTicketInfo(order uuid.UUID) error {
	return nil
}
