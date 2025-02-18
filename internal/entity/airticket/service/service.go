package service

import (
	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type TicketHandler interface {
	GetAll() ([]entity.AirTicket, error)
	Edit(order uuid.UUID, edited entity.AirTicket) error
	Remove(order uuid.UUID) error
}

type AirticketService struct {
	ticketHandler TicketHandler
}

func New(ticketHandler TicketHandler) *AirticketService {
	return &AirticketService{
		ticketHandler: ticketHandler,
	}
}
