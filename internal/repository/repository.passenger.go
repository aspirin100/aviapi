package repository

import (
	"context"
	"fmt"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

func (repo *Repository) GetPassengerList(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.Passenger, error) {
	ex := repo.CheckTx(ctx)

	passengers := []entity.Passenger{}

	err := ex.SelectContext(
		ctx,
		&passengers,
		GetPassengerListQuery,
		ticketOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get passengers list: %w", err)
	}

	return passengers, nil
}

func (repo *Repository) EditPassengerInfo(
	ctx context.Context,
	passengerID uuid.UUID,
	edited entity.Passenger) (*entity.Passenger, error) {

	return nil, nil
}

func (repo *Repository) RemovePassengerInfo(passengerID uuid.UUID) error {
	return nil
}

const (
	GetPassengerListQuery = `
	SELECT
		first_name,
		last_name,
		patronymic
	FROM ticket_passengers
	JOIN passengers
	ON passengers.id = ticket_passengers.passenger_id
	WHERE order_id = $1;
	`
)
