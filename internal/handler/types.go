package handler

import (
	"context"
	"time"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

type TicketManager interface {
	GetTicketList(ctx context.Context) ([]entity.AirTicket, error)
	EditTicket(ctx context.Context, order uuid.UUID, edited entity.AirTicket) (*entity.AirTicket, error)
	RemoveTicketInfo(ctx context.Context, order uuid.UUID) error
}

type PassengerManager interface {
	GetPassengerList(ticketOrderID uuid.UUID) ([]entity.Passenger, error)
	EditPassengerInfo(passengerID uuid.UUID, edited entity.Passenger) error
	RemovePassengerInfo(passengerID uuid.UUID) error
}

type DocumentManager interface {
	GetDocumentList(ctx context.Context, passengerID uuid.UUID) ([]entity.Document, error)
	EditDocumentInfo(ctx context.Context, documentID uuid.UUID, edited entity.Document) (*entity.Document, error)
	RemoveDocumentInfo(documentID uuid.UUID) error
}

type AirflightManager interface {
	TicketManager
	PassengerManager
	DocumentManager
	GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error)
	GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error)
}
