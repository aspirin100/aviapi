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
	Type string    `db:"document_type" json:"document_type"`
	ID   uuid.UUID `db:"id" json:"id,omitempty"`
}

type PassengerWithDocuments struct {
	PassengerID uuid.UUID  `db:"passenger_id" json:"passenger_id"`
	FirstName   string     `db:"first_name" json:"first_name"`
	LastName    string     `db:"last_name" json:"last_name"`
	Patronymic  string     `db:"patronymic" json:"patronymic"`
	Documents   []Document `db:"documents" json:"documents"`
}

type FullInfo struct {
	OrderID          uuid.UUID                `db:"order_id" json:"order_id"`
	FromCountry      string                   `db:"from_country" json:"from_country"`
	ToCountry        string                   `db:"to_country" json:"to_country"`
	Carrier          string                   `db:"carrier" json:"carrier"`
	DepartureDate    time.Time                `db:"departure_date" json:"departure_date"`
	ArrivalDate      time.Time                `db:"arrival_date" json:"arrival_date"`
	RegistrationDate time.Time                `db:"registration_date" json:"registration_date"`
	Passengers       []PassengerWithDocuments `json:"passengers"`
}

type Report struct {
	RegistrationDate time.Time  `db:"registration_date" json:"registration_date"`
	DepartureDate    *time.Time `db:"departure_date" json:"departure_date"`
	OrderID          uuid.UUID  `db:"order_id" json:"order_id"`
	FromCountry      string     `db:"from_country" json:"from"`
	ToCountry        string     `db:"to_country" json:"to"`
	ServiceProvided  bool       `db:"service_provided" json:"provided"`
}
