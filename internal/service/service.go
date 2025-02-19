package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type AirflightHandler interface {
	GetFullInfo(ctx context.Context, ticketOrderID uuid.UUID) (*entity.FullInfo, error)
	GetReport(ctx context.Context,
		passengerID uuid.UUID,
		periodStart, periodEnd time.Time) ([]entity.Report, error)
	BeginTx(ctx context.Context) (context.Context, entity.CommitOrRollback, error)
}

type InfoService struct {
	airflightHandler AirflightHandler
}

func NewInfoService(airflightHandler AirflightHandler) *InfoService {
	return &InfoService{
		airflightHandler: airflightHandler,
	}
}

func (s *InfoService) GetFullInfo(ctx context.Context, ticketOrderID uuid.UUID) (*entity.FullInfo, error) {
	const op = "service.GetFullInfo"

	ctx, commitOrRollback, err := s.airflightHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	fullinfo, err := s.airflightHandler.GetFullInfo(ctx, ticketOrderID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrTicketNotFound):
			return nil, entity.ErrTicketNotFound
		default:
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return fullinfo, nil
}

func (s *InfoService) GetReport(ctx context.Context,
	passengerID uuid.UUID,
	periodStart,
	periodEnd time.Time) ([]entity.Report, error) {
	const op = "service.GetReport"

	ctx, commitOrRollback, err := s.airflightHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	report, err := s.airflightHandler.GetReport(
		ctx,
		passengerID,
		periodStart,
		periodEnd)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return report, nil
}
