package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

const (
	AveragePassengerCount = 200
)

func (repo *Repository) GetTicketList(ctx context.Context) ([]entity.AirTicket, error) {
	ex := repo.CheckTx(ctx)

	tickets := make([]entity.AirTicket, 0, AveragePassengerCount)

	err := ex.SelectContext(ctx, &tickets, GetTicketListQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get tickets list: %w", err)
	}

	//memory usage optimisation
	finalList := make([]entity.AirTicket, len(tickets))
	copy(finalList, tickets)

	return finalList, nil
}

func (repo *Repository) EditTicket(order uuid.UUID, edited entity.AirTicket) error {
	return nil
}

func (repo *Repository) RemoveTicketInfo(order uuid.UUID) error {
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
)
