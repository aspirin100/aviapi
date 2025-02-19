package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/aspirin100/aviapi/internal/repository"
)

type AirflightHandler interface {
	GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error)
	GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error)
	BeginTx(ctx context.Context) (context.Context, repository.CommitOrRollback, error)
}

type InfoService struct {
	airflightHandler AirflightHandler
}

func NewInfoService(airflightHandler AirflightHandler) *InfoService {
	return &InfoService{
		airflightHandler: airflightHandler,
	}
}

func (s *InfoService) GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error) {
	return nil, nil
}

func (s *InfoService) GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error) {
	return nil, nil
}
