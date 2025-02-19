package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type TicketHandler interface {
	GetTicketList(ctx context.Context) ([]entity.AirTicket, error)
	EditTicket(ctx context.Context, order uuid.UUID, edited entity.AirTicket) (*entity.AirTicket, error)
	RemoveTicketInfo(ctx context.Context, order uuid.UUID) error
	BeginTx(ctx context.Context) (context.Context, entity.CommitOrRollback, error)
}

type AirticketService struct {
	ticketHandler TicketHandler
}

func NewAirticketService(ticketHandler TicketHandler) *AirticketService {
	return &AirticketService{
		ticketHandler: ticketHandler,
	}
}

func (as *AirticketService) GetTicketList(ctx context.Context) ([]entity.AirTicket, error) {
	const op = "service.GetTicketList"

	ctx, commitOrRollback, err := as.ticketHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	tickets, err := as.ticketHandler.GetTicketList(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tickets, nil
}

func (as *AirticketService) EditTicket(
	ctx context.Context,
	order uuid.UUID,
	edited entity.AirTicket) (*entity.AirTicket, error) {
	ctx, commitOrRollback, err := as.ticketHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	changedTicketInfo, err := as.ticketHandler.EditTicket(ctx, order, edited)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrTicketNotFound):
			return nil, entity.ErrTicketNotFound
		default:
			return nil, fmt.Errorf("failed to edit ticket: %w", err)
		}
	}

	return changedTicketInfo, nil
}

func (as *AirticketService) RemoveTicketInfo(
	ctx context.Context,
	order uuid.UUID) error {
	const op = "service.RemoveTicketInfo"

	ctx, commitOrRollback, err := as.ticketHandler.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	err = as.ticketHandler.RemoveTicketInfo(ctx, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
