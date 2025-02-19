package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/aspirin100/aviapi/internal/repository"
	"github.com/aspirin100/aviapi/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	PostgresDSN = "postgres://postgres:postgres@localhost:5432/aviapi_db?sslmode=disable"
)

var (
	uids = []uuid.UUID{
		uuid.MustParse("99ef07fa-a5b9-4c14-b409-c4ca83ff16a6"),
		uuid.MustParse("e5c0536c-a544-4cdd-b3cb-a9c31eaac78f"),
		uuid.MustParse("9b3b2ed9-e21f-4493-b550-76595cf37785"),
	}
)

func initAirticetService() (*service.AirticketService, error) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
	if err != nil {
		return nil, err
	}

	return service.NewAirticketService(repo), nil
}

func TestEditTicketInfo(t *testing.T) {
	srv, err := initAirticetService()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	type Params struct {
		orderId uuid.UUID
		edited  entity.AirTicket
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
				orderId: uids[0],
				edited: entity.AirTicket{
					From: "Edited deportation country",
				},
			},
		},
		{
			Name:        "not found case",
			ExpectedErr: service.ErrTicketNotFound,
			Args: Params{
				orderId: uuid.Nil,
				edited: entity.AirTicket{
					From: "Edited deportation country",
				},
			},
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			_, err := srv.EditTicket(
				context.Background(),
				tcase.Args.orderId,
				tcase.Args.edited)
			
			require.EqualValues(t, tcase.ExpectedErr, err)
		})
	}
}
