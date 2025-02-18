package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type TicketHandler interface {
	GetTicketList(ctx context.Context) ([]entity.AirTicket, error)
	EditTicket(ctx context.Context, order uuid.UUID, edited entity.AirTicket) (*entity.AirTicket, error)
	RemoveTicketInfo(ctx context.Context, order uuid.UUID) error
}

type AirticketService struct {
	ticketHandler TicketHandler
}

func New(ticketHandler TicketHandler) *AirticketService {
	return &AirticketService{
		ticketHandler: ticketHandler,
	}
}

func (as *AirticketService) GetTicketList(ctx context.Context) ([]entity.AirTicket, error) {
	tickets, err := as.ticketHandler.GetTicketList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket list: %w", err)
	}

	return tickets, nil
}

func (as *AirticketService) EditTicket(ctx context.Context, order uuid.UUID, edited entity.AirTicket) (*entity.AirTicket, error) {
	return nil, nil
}

func (as *AirticketService) RemoveTicketInfo(ctx context.Context, order uuid.UUID) error {
	return nil
}
