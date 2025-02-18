package entity

import (
	"time"

	"github.com/google/uuid"
)

type AirTicket struct {
	From             string    `db:"from"`
	To               string    `db:"to"`
	Order            uuid.UUID `db:"order_id"`
	Carrier          string    `db:"carrier"`
	DepartureDate    time.Time `db:"departure_date"`
	ArrivalDate      time.Time `db:"arrival_date"`
	RegistrationDate time.Time `db:"registration_date"`
}

type Passenger struct {
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Patronymic string `db:"patronymic"`
}

type Document struct {
	Type string    `db:"type"`
	ID   uuid.UUID `db:"id"`
}
