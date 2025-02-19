package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

func (repo *Repository) GetTicketList(ctx context.Context) ([]entity.AirTicket, error) {
	ex := repo.CheckTx(ctx)

	tickets := []entity.AirTicket{}

	err := ex.SelectContext(ctx, &tickets, GetTicketListQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get tickets list: %w", err)
	}

	return tickets, nil
}

func (repo *Repository) EditTicketInfo(
	ctx context.Context,
	order uuid.UUID,
	edited *entity.AirTicket) (*entity.AirTicket, error) {
	ex := repo.CheckTx(ctx)

	var changedTicket entity.AirTicket

	err := ex.GetContext(ctx, &changedTicket, EditTicketInfoQuery,
		order,
		edited.From,
		edited.To,
		edited.Carrier,
		edited.DepartureDate,
		edited.ArrivalDate,
		edited.RegistrationDate,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, entity.ErrTicketNotFound
		default:
			return nil, fmt.Errorf("failed to edit info: %w", err)
		}
	}

	return &changedTicket, nil
}

func (repo *Repository) RemoveTicketInfo(ctx context.Context, order uuid.UUID) error {
	ex := repo.CheckTx(ctx)

	//nolint:gocritic
	_, err := ex.QueryContext(
		ctx,
		RemoveTicketInfoQuery,
		order)
	if err != nil {
		return fmt.Errorf("failed to remove ticket info: %w", err)
	}

	return nil
}

const (
	GetTicketListQuery = `
	SELECT 
		order_id,
		from_country,
		to_country,
		carrier,
		departure_date,
		arrival_date,
		registration_date
	FROM tickets;
	`

	EditTicketInfoQuery = `
    UPDATE tickets SET
        from_country = CASE WHEN $2 = '' THEN from_country ELSE $2 END,
        to_country = CASE WHEN $3 = '' THEN to_country ELSE $3 END,
        carrier = CASE WHEN $4 = '' THEN carrier ELSE $4 END,
        departure_date = CASE WHEN $5::TIMESTAMPTZ IS NULL THEN departure_date ELSE $5::TIMESTAMPTZ END,
        arrival_date = CASE WHEN $6::TIMESTAMPTZ IS NULL THEN arrival_date ELSE $6::TIMESTAMPTZ END,
        registration_date = CASE WHEN $7::TIMESTAMPTZ IS NULL THEN registration_date ELSE $7::TIMESTAMPTZ END
    WHERE
        order_id = $1
    RETURNING 
        order_id,
        from_country,
        to_country,
        carrier,
        departure_date,
        arrival_date,
        registration_date;
	`

	RemoveTicketInfoQuery = `
	DELETE FROM tickets
	WHERE order_id = $1;
	`
)
