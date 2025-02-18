package service

import (
	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type TicketHandler interface {
	GetTicketList() ([]entity.AirTicket, error)
	EditTicket(order uuid.UUID, edited entity.AirTicket) error
	RemoveTicketInfo(order uuid.UUID) error
}

type AirticketService struct {
	ticketHandler TicketHandler
}

func New(ticketHandler TicketHandler) *AirticketService {
	return &AirticketService{
		ticketHandler: ticketHandler,
	}
}

func (as *AirticketService) GetTicketList() ([]entity.AirTicket, error) {
	return nil, nil
}

func (as *AirticketService) EditTicket(order uuid.UUID, edited entity.AirTicket) error {
	return nil
}

func (as *AirticketService) RemoveTicketInfo(order uuid.UUID) error {
	return nil
}
