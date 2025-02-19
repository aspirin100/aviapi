package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type AirflightHandler interface {
	GetFullInfo(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.FullInfo, error)
	GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error)
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

func (s *InfoService) GetFullInfo(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.FullInfo, error) {
	return nil, nil
}

func (s *InfoService) GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error) {
	return nil, nil
}
