package service

import (
	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

type PassengerHandler interface {
	GetPassengerList(ticketOrderID uuid.UUID) ([]entity.Passenger, error)
	EditPassengerInfo(passengerID uuid.UUID, edited entity.Passenger) error
	RemovePassengerInfo(passengerID uuid.UUID) error
}

type PassengerService struct {
	passengerHandler PassengerHandler
}

func New(passengerHandler PassengerHandler) *PassengerService {
	return &PassengerService{
		passengerHandler: passengerHandler,
	}
}

func (ps *PassengerService) GetPassengerList(ticketOrderID uuid.UUID) ([]entity.Passenger, error) {
	return nil, nil
}

func (ps *PassengerService) EditPassengerInfo(passengerID uuid.UUID, edited entity.Passenger) error {
	return nil
}

func (ps *PassengerService) RemovePassengerInfo(passengerID uuid.UUID) error {
	return nil
}
