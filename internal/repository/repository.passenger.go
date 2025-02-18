package repository

import (
	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

func (repo *Repository) GitPassengerList(ticketOrderID uuid.UUID) ([]entity.Passenger, error) {
	return nil, nil
}

func (repo *Repository) EditPassengerInfo(passengerID uuid.UUID, edited entity.Passenger) error {
	return nil
}

func (repo *Repository) RemovePassengerInfo(passengerID uuid.UUID) error {
	return nil
}
