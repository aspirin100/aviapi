package service

import (
	"context"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

type PassengerHandler interface {
	GetPassengerList(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.Passenger, error)
	EditPassengerInfo(ctx context.Context, passengerID uuid.UUID, edited entity.Passenger) (*entity.Passenger, error)
	RemovePassengerInfo(ctx context.Context, passengerID uuid.UUID) error
}

type PassengerService struct {
	passengerHandler PassengerHandler
}

func New(passengerHandler PassengerHandler) *PassengerService {
	return &PassengerService{
		passengerHandler: passengerHandler,
	}
}

func (ps *PassengerService) GetPassengerList(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.Passenger, error) {
	return nil, nil
}

func (ps *PassengerService) EditPassengerInfo(ctx context.Context, passengerID uuid.UUID, edited entity.Passenger) (*entity.Passenger, error) {
	return nil, nil
}

func (ps *PassengerService) RemovePassengerInfo(ctx context.Context, passengerID uuid.UUID) error {
	return nil
}
