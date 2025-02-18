package repository

import (
	"fmt"
	"time"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sqlx.DB
}

func NewConnection(driverName, DSN string) (*Repository, error) {
	db, err := sqlx.Connect(driverName, DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Repository{
		DB: db,
	}, nil
}

func (repo *Repository) GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error) {
	return nil, nil
}

func (repo *Repository) GetReport(
	passengerID uuid.UUID,
	periodStart, periodEnd time.Time) ([]entity.AirTicket, error) {
	return nil, nil
}
