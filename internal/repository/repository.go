package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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

func (repo *Repository) GetFullInfo(ctx context.Context, ticketOrderID uuid.UUID) (*entity.FullInfo, error) {
	ex := repo.CheckTx(ctx)

	var rows []struct {
		OrderID          uuid.UUID `db:"order_id"`
		FromCountry      string    `db:"from_country"`
		ToCountry        string    `db:"to_country"`
		Carrier          string    `db:"carrier"`
		DepartureDate    time.Time `db:"departure_date"`
		ArrivalDate      time.Time `db:"arrival_date"`
		RegistrationDate time.Time `db:"registration_date"`
		PassengerID      uuid.UUID `db:"passenger_id"`
		FirstName        string    `db:"first_name"`
		LastName         string    `db:"last_name"`
		Patronymic       string    `db:"patronymic"`
		Documents        []byte    `db:"documents"`
	}

	err := ex.SelectContext(
		ctx,
		&rows,
		GetFullInfoQuery,
		ticketOrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to read all info: %w", err)
	}

	if len(rows) == 0 {
		return nil, entity.ErrTicketNotFound
	}

	info := entity.FullInfo{
		OrderID:          rows[0].OrderID,
		FromCountry:      rows[0].FromCountry,
		ToCountry:        rows[0].ToCountry,
		Carrier:          rows[0].Carrier,
		DepartureDate:    rows[0].DepartureDate,
		ArrivalDate:      rows[0].ArrivalDate,
		RegistrationDate: rows[0].RegistrationDate,
		Passengers:       []entity.PassengerWithDocuments{},
	}

	for _, row := range rows {
		var docs []entity.Document

		if len(row.Documents) > 0 {
			json.Unmarshal(row.Documents, &docs)
		}

		info.Passengers = append(
			info.Passengers,
			entity.PassengerWithDocuments{
				PassengerID: row.PassengerID,
				FirstName:   row.FirstName,
				LastName:    row.LastName,
				Patronymic:  row.Patronymic,
				Documents:   docs,
			})
	}

	return &info, nil
}

func (repo *Repository) GetReport(ctx context.Context,
	passengerID uuid.UUID,
	periodStart, periodEnd time.Time) ([]entity.Report, error) {
	ex := repo.CheckTx(ctx)

	report := []entity.Report{}

	err := ex.SelectContext(ctx, &report, GetReportQuery,
		periodStart,
		periodEnd,
		passengerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get report: %w", err)
	}

	return report, nil
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
		p.id AS passenger_id,
		p.first_name,
		p.last_name,
		COALESCE(p.patronymic, '') AS patronymic,
		JSON_AGG(
			JSON_BUILD_OBJECT(
				'id', d.id,
				'document_type', d.document_type
			)
		) FILTER (WHERE d.id IS NOT NULL) AS documents
	FROM tickets t
	LEFT JOIN ticket_passengers tp ON t.order_id = tp.order_id
	LEFT JOIN passengers p ON tp.passenger_id = p.id
	LEFT JOIN documents d ON p.id = d.passenger_id
	WHERE t.order_id = $1
	GROUP BY t.order_id, p.id
		`
	GetReportQuery = `
	SELECT
		t.registration_date,
		t.departure_date,
		t.order_id,
		t.from_country,
		t.to_country,
		CASE
			WHEN t.departure_date <= $2 THEN TRUE
			ELSE FALSE
		END AS service_provided
	FROM tickets t
	LEFT JOIN ticket_passengers tp ON t.order_id = tp.order_id
	LEFT JOIN passengers p ON tp.passenger_id = p.id
	WHERE
		p.id = $3
		AND (
			(t.registration_date < $1 AND t.departure_date >= $1 AND t.departure_date <= $2)
			OR
			(t.registration_date >= $1 AND t.registration_date <= $2 AND (t.departure_date > $2 OR t.departure_date IS NULL))
			OR
			(t.registration_date >= $1 AND t.registration_date <= $2 AND t.departure_date >= $1 AND t.departure_date <= $2)
		)
	ORDER BY t.registration_date;
	`
)
