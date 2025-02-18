package handler

import (
	"time"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

type TicketManager interface {
	GetTicketList() ([]entity.AirTicket, error)
	EditTicket(order uuid.UUID, edited entity.AirTicket) error
	RemoveTicketInfo(order uuid.UUID) error
}

type PassengerManager interface {
	GitPassengerList(ticketOrderID uuid.UUID) ([]entity.Passenger, error)
	EditPassengerInfo(passengerID uuid.UUID, edited entity.Passenger) error
	RemovePassengerInfo(passengerID uuid.UUID) error
}

type DocumentManager interface {
	GetDocumentList(passengerID uuid.UUID) ([]entity.Document, error)
	EditDocumentInfo(documentID uuid.UUID, edited entity.Document) error
	RemoveDocumentInfo(documentID uuid.UUID) error
}

type AirflightManager interface {
	TicketManager
	PassengerManager
	DocumentManager
	GetFullInfo(ticketOrderID uuid.UUID) (*entity.FullInfo, error)
	GetReport(passengerID uuid.UUID, periodStart, periodEnd time.Time) ([]entity.AirTicket, error)
}
