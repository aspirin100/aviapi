package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/aspirin100/aviapi/internal/repository"
)

const (
	PostgresDSN = "postgres://postgres:postgres@localhost:5432/aviapi_db?sslmode=disable"
)

// only for test, changes with every migration.
var (
	TicketIDs = []uuid.UUID{
		uuid.MustParse("af4ca61e-9810-48f6-aceb-7733713ca7c9"),
	}
)

func TestGetTicketList(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
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

func TestEditTicketInfo(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
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
				orderID: TicketIDs[1],
				edited: entity.AirTicket{
					From: "EDITED",
					To:   "EDITED",
				},
			},
		},
		{
			Name:        "ticket not found case",
			ExpectedErr: entity.ErrTicketNotFound,
			Args: Params{
				orderID: uuid.Nil,
				edited: entity.AirTicket{
					From: "EDITED",
					To:   "EDITED",
				},
			},
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			_, err := repo.EditTicketInfo(context.Background(),
				tcase.Args.orderID,
				&tcase.Args.edited)

			require.EqualValues(t, tcase.ExpectedErr, err)
		})
	}
}

func TestRemoveTicketInfo(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
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
			OrderID:     TicketIDs[0],
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
