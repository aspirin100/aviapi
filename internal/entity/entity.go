package entity

import (
	"time"

	"github.com/google/uuid"
)

type AirTicket struct {
	From             string     `db:"from_country" json:"from"`
	To               string     `db:"to_country" json:"to"`
	Order            uuid.UUID  `db:"order_id" json:"order_id"`
	Carrier          string     `db:"carrier" json:"carrier"`
	DepartureDate    *time.Time `db:"departure_date" json:"departure_date"`
	ArrivalDate      *time.Time `db:"arrival_date" json:"arrival_date"`
	RegistrationDate *time.Time `db:"registration_date" json:"registration_date"`
}

type Passenger struct {
	FirstName  string `db:"first_name" json:"first_name"`
	LastName   string `db:"last_name" json:"last_name"`
	Patronymic string `db:"patronymic" json:"patronymic"`
}

type Document struct {
	Type string    `db:"type" json:"type"`
	ID   uuid.UUID `db:"id" json:"id"`
}

type FullInfo struct {
	Ticket        AirTicket
	PassengerList []Passenger
	DocumentList  []Document
}
