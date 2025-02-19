package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CommitOrRollback func(err error) error

var (
	ErrDocumentNotFound  = errors.New("document not found")
	ErrTicketNotFound    = errors.New("ticket not found")
	ErrPassengerNotFound = errors.New("passenger not found")
)

type AirTicket struct {
	From             string     `db:"from_country" json:"from,omitempty"`
	To               string     `db:"to_country" json:"to,omitempty"`
	Order            uuid.UUID  `db:"order_id" json:"order_id,omitempty"`
	Carrier          string     `db:"carrier" json:"carrier,omitempty"`
	DepartureDate    *time.Time `db:"departure_date" json:"departure_date,omitempty"`
	ArrivalDate      *time.Time `db:"arrival_date" json:"arrival_date,omitempty"`
	RegistrationDate *time.Time `db:"registration_date" json:"registration_date,omitempty"`
}

type Passenger struct {
	FirstName  string `db:"first_name" json:"first_name,omitempty"`
	LastName   string `db:"last_name" json:"last_name,omitempty"`
	Patronymic string `db:"patronymic" json:"patronymic,omitempty"`
}

type Document struct {
	Type string    `db:"document_type" json:"type"`
	ID   uuid.UUID `db:"id" json:"id,omitempty"`
}

type FullInfo struct {
	From             string     `db:"from_country" json:"from"`
	To               string     `db:"to_country" json:"to"`
	Order            uuid.UUID  `db:"order_id" json:"order_id"`
	Carrier          string     `db:"carrier" json:"carrier"`
	DepartureDate    *time.Time `db:"departure_date" json:"departure_date"`
	ArrivalDate      *time.Time `db:"arrival_date" json:"arrival_date"`
	RegistrationDate *time.Time `db:"registration_date" json:"registration_date"`
	FirstName        string     `db:"first_name" json:"first_name"`
	LastName         string     `db:"last_name" json:"last_name"`
	Patronymic       string     `db:"patronymic" json:"patronymic,omitempty"`
	DocumentType     string     `db:"document_type" json:"document_type"`
	DocumentID       uuid.UUID  `db:"document_id" json:"document_id,omitempty"`
	PassengerID      uuid.UUID  `db:"passenger_id" json:"passenger_id"`
}
