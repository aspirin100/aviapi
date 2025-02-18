package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/aspirin100/aviapi/internal/entity"
)

const (
	PostgresDSN = "postgres://postgres:postgres@localhost:5432/aviapi_db?sslmode=disable"
)

// only for test, changes with every migration
var (
	ticketIDs = []uuid.UUID{
		uuid.MustParse("2eda32fd-0c41-4654-bb2e-5d45760e02a8"),
		uuid.MustParse("62973252-1ad9-447e-ac7a-39a8daada566"),
		uuid.MustParse("ac2057bc-8fdc-47ae-af72-fb36c65f7abf"),
	}
)

func TestGetTicketList(t *testing.T) {
	repo, err := NewConnection("postgres", PostgresDSN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	list, err := repo.GetTicketList(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, val := range list {
		fmt.Println(val)
	}
}

func TestEditTicket(t *testing.T) {
	repo, err := NewConnection("postgres", PostgresDSN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	type Params struct {
		orderID uuid.UUID
		edited  entity.AirTicket
	}

	cases := []struct {
		Name        string
		ExpectedErr error
		Args        Params
	}{
		{
			Name:        "partitial edit",
			ExpectedErr: nil,
			Args: Params{
				orderID: ticketIDs[1],
				edited: entity.AirTicket{
					From: "EDITED",
					To:   "EDITED",
				},
			},
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			_, err := repo.EditTicket(context.Background(), cases[0].Args.orderID, cases[0].Args.edited)

			require.EqualValues(t, tcase.ExpectedErr, err)
		})
	}

}

func TestRemoveTicketInfo(t *testing.T) {
	repo, err := NewConnection("postgres", PostgresDSN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cases := []struct {
		Name        string
		ExpectedErr error
		OrderID     uuid.UUID
	}{
		{
			Name:        "ok case",
			ExpectedErr: nil,
			OrderID:     ticketIDs[0],
		},
		{
			Name:        "ticket not found case",
			ExpectedErr: nil,
			OrderID:     uuid.Nil,
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			err := repo.RemoveTicketInfo(context.Background(), tcase.OrderID)

			require.EqualValues(t, tcase.ExpectedErr, err)
		})
	}
}
