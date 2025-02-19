package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aspirin100/aviapi/internal/repository"
)

func TestGetFullInfo(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fullinfo, err := repo.GetFullInfo(context.Background(), TicketIDs[0])
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(fullinfo)
}
