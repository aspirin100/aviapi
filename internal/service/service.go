package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type AirflightHandler interface {
	GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error)
	GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error)
}

type Service struct {
	airflightHandler AirflightHandler
}

func New(airflightHandler AirflightHandler) *Service {
	return &Service{
		airflightHandler: airflightHandler,
	}
}

func (s *Service) GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error) {
	return nil, nil
}

func (s *Service) GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error) {
	return nil, nil
}
