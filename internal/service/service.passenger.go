package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type PassengerHandler interface {
	GetPassengerList(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.Passenger, error)
	EditPassengerInfo(ctx context.Context, passengerID uuid.UUID, edited entity.Passenger) (*entity.Passenger, error)
	RemovePassengerInfo(ctx context.Context, passengerID uuid.UUID) error
	BeginTx(ctx context.Context) (context.Context, entity.CommitOrRollback, error)
}

type PassengerService struct {
	passengerHandler PassengerHandler
}

func NewPassengerService(passengerHandler PassengerHandler) *PassengerService {
	return &PassengerService{
		passengerHandler: passengerHandler,
	}
}

func (ps *PassengerService) GetPassengerList(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.Passenger, error) {
	const op = "service.GetPassengersList"

	ctx, commitOrRollback, err := ps.passengerHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	passengers, err := ps.passengerHandler.GetPassengerList(
		ctx,
		ticketOrderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return passengers, nil
}

func (ps *PassengerService) EditPassengerInfo(
	ctx context.Context,
	passengerID uuid.UUID,
	edited entity.Passenger) (*entity.Passenger, error) {
	ctx, commitOrRollback, err := ps.passengerHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	changedPassengerInfo, err := ps.passengerHandler.EditPassengerInfo(
		ctx,
		passengerID,
		edited)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrPassengerNotFound):
			return nil, entity.ErrPassengerNotFound
		default:
			return nil, fmt.Errorf("failed to edit passenger info: %w", err)
		}
	}

	return changedPassengerInfo, nil
}

func (ps *PassengerService) RemovePassengerInfo(
	ctx context.Context,
	passengerID uuid.UUID) error {
	const op = "service.RemovePassengerInfo"

	ctx, commitOrRollback, err := ps.passengerHandler.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	err = ps.passengerHandler.RemovePassengerInfo(
		ctx,
		passengerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
