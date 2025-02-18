package repository

import (
	"context"
	"fmt"
	"testing"
)

const (
	PostgresDSN = "postgres://postgres:postgres@localhost:5432/aviapi_db?sslmode=disable"
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
