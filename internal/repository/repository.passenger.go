package repository

import (
	"context"
	"database/sql"
	"errors"
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
	ex := repo.CheckTx(ctx)

	var changedPassengerInfo entity.Passenger

	err := ex.GetContext(ctx, &changedPassengerInfo, EditPassengerInfoQuery,
		passengerID,
		edited.FirstName,
		edited.LastName,
		edited.Patronymic,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, entity.ErrPassengerNotFound
		default:
			return nil, fmt.Errorf("failed to edit passenger info: %w", err)
		}
	}

	return &changedPassengerInfo, nil
}

func (repo *Repository) RemovePassengerInfo(ctx context.Context, passengerID uuid.UUID) error {
	ex := repo.CheckTx(ctx)

	//nolint:gocritic
	_, err := ex.QueryContext( 
		ctx,
		RemovePassengerInfoQuery,
		passengerID)
	if err != nil {
		return fmt.Errorf("failed to remove passenger info: %w", err)
	}

	return nil
}

const (
	GetPassengerListQuery = `
	SELECT
		first_name,
		last_name,
		COALESCE(patronymic, '') AS patronymic
	FROM ticket_passengers
	JOIN passengers
	ON passengers.id = ticket_passengers.passenger_id
	WHERE order_id = $1;
	`

	EditPassengerInfoQuery = `
	UPDATE passengers SET
		first_name = CASE WHEN $2 = '' THEN first_name ELSE $2 END,
		last_name = CASE WHEN $3 = '' THEN last_name ELSE $3 END,
		patronymic = CASE WHEN $4 = '' THEN patronymic ELSE $4 END
	WHERE
		id = $1
	RETURNING 
		first_name,
		last_name,
		COALESCE(patronymic, '') AS patronymic;
	`

	RemovePassengerInfoQuery = `
	DELETE FROM passengers
	WHERE id = $1;
	`
)
