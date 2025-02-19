package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/aspirin100/aviapi/internal/entity"
)

type Repository struct {
	DB *sqlx.DB
}

type executor interface {
	QueryContext(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
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

type ctxKey struct{}

var txContextKey = ctxKey{}

func (r *Repository) BeginTx(ctx context.Context) (context.Context, entity.CommitOrRollback, error) {
	tx, err := r.DB.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	return context.WithValue(ctx, txContextKey, tx), func(err error) error {
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				return errors.Join(err, errRollback)
			}

			return err
		}

		errCommit := tx.Commit()
		if errCommit != nil {
			return fmt.Errorf("failed to commit transaction: %w", errCommit)
		}

		return nil
	}, nil
}

func (r *Repository) CheckTx(ctx context.Context) executor {
	var ex executor = r.DB

	// checks if current operation is in transaction
	tx, ok := ctx.Value(txContextKey).(sqlx.Tx)
	if ok {
		ex = &tx
	}

	return ex
}

func (repo *Repository) GetFullInfo(
	ctx context.Context,
	ticketOrderID uuid.UUID) ([]entity.FullInfo, error) {
	ex := repo.CheckTx(ctx)

	info := []entity.FullInfo{}

	err := ex.SelectContext(
		ctx,
		&info,
		GetFullInfoQuery,
		ticketOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to read all info: %w", err)
	}

	return info, nil
}

func (repo *Repository) GetReport(
	passengerID uuid.UUID,
	periodStart, periodEnd time.Time) ([]entity.AirTicket, error) {
	return nil, nil
}

const (
	GetFullInfoQuery = `
	SELECT
		t.order_id,
		t.from_country,
		t.to_country,
		t.carrier,
		t.departure_date,
		t.arrival_date,
		t.registration_date,
		p.first_name,
		p.last_name,
		COALESCE(p.patronymic, '') AS patronymic,
		d.id AS document_id,
		d.document_type,
		d.passenger_id
	FROM tickets t
	LEFT JOIN ticket_passengers tp ON t.order_id = tp.order_id
	LEFT JOIN passengers p ON tp.passenger_id = p.id
	LEFT JOIN documents d ON p.id = d.passenger_id
	WHERE t.order_id = $1;
	`
)
