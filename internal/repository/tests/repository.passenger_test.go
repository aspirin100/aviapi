package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/aspirin100/aviapi/internal/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	pasIDs = []uuid.UUID{
		uuid.MustParse("368a53bd-db1b-457e-bb21-764bfcc895bb"),
		uuid.MustParse("3da7e550-cd8a-4824-a876-1e337b3eec60"),
		uuid.MustParse("67e320c8-b3c6-46f1-8553-f4e98696e2b5"),
	}
)

func TestGetPassengerList(t *testing.T) {
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
			OrderID:     ticketIDs[2],
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			passengers, err := repo.GetPassengerList(context.Background(), tcase.OrderID)

			require.EqualValues(t, tcase.ExpectedErr, err)

			fmt.Println(passengers)
		})
	}
}

func TestEditPassengerInfo(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	type Params struct {
		passengerID uuid.UUID
		edited      entity.Passenger
	}

	cases := []struct {
		Name        string
		ExpectedErr error
		Args        Params
	}{
		{
			Name:        "ok case",
			ExpectedErr: nil,
			Args: Params{
				passengerID: pasIDs[0],
				edited: entity.Passenger{
					FirstName: "EDITED_FIRST_NAME",
				},
			},
		},
		{
			Name:        "passenger not found case",
			ExpectedErr: repository.ErrPassengerNotFound,
			Args: Params{
				passengerID: uuid.Nil,
				edited: entity.Passenger{
					FirstName: "EDITED_FIRST_NAME",
				},
			},
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			passengers, err := repo.EditPassengerInfo(
				context.Background(),
				tcase.Args.passengerID,
				tcase.Args.edited)

			require.EqualValues(t, tcase.ExpectedErr, err)

			fmt.Println(passengers)
		})
	}
}
