package handler

import (
	"context"
	"time"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

type TicketManager interface {
	GetTicketList(ctx context.Context) ([]entity.AirTicket, error)
	EditTicketInfo(ctx context.Context,
		order uuid.UUID,
		edited *entity.AirTicket) (*entity.AirTicket, error)
	RemoveTicketInfo(ctx context.Context, order uuid.UUID) error
}

type PassengerManager interface {
	GetPassengerList(ctx context.Context, ticketOrderID uuid.UUID) ([]entity.Passenger, error)
	EditPassengerInfo(ctx context.Context,
		passengerID uuid.UUID,
		edited entity.Passenger) (*entity.Passenger, error)
	RemovePassengerInfo(ctx context.Context, passengerID uuid.UUID) error
}

type DocumentManager interface {
	GetDocumentList(ctx context.Context, passengerID uuid.UUID) ([]entity.Document, error)
	EditDocumentInfo(ctx context.Context,
		documentID uuid.UUID,
		edited entity.Document) (*entity.Document, error)
	RemoveDocumentInfo(ctx context.Context, documentID uuid.UUID) error
}

type AirflightManager interface {
	TicketManager
	PassengerManager
	DocumentManager
	GetFullInfo(ctx context.Context, ticketOrderID uuid.UUID) (*entity.FullInfo, error)
	GetReport(ctx context.Context,
		passengerID uuid.UUID,
		periodStart,
		periodEnd time.Time) ([]entity.Report, error)
}
