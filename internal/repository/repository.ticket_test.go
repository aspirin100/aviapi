package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

const (
	PostgresDSN = "postgres://postgres:postgres@localhost:5432/aviapi_db?sslmode=disable"
)

var (
	ticketIDs = []uuid.UUID{
		uuid.MustParse("ffa6a9a4-6e0f-4bbb-81c6-e07f7eb2547b"),
		uuid.MustParse("82a20871-3146-4d10-a8bd-5f0f30ca18d0"),
		uuid.MustParse("8b8c1779-de07-43cb-85a9-613550971d45"),
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

	edited, err := repo.EditTicket(context.Background(), cases[0].Args.orderID, cases[0].Args.edited)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(edited)
}
