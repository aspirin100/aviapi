package service

import (
	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

type PassengerHandler interface {
	GitPassengerList(ticketOrderID uuid.UUID) ([]entity.Passenger, error)
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
