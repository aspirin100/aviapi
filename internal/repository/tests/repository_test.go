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

	infoList, err := repo.GetFullInfo(context.Background(), TicketIDs[0])
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, val := range infoList {
		fmt.Print(val, "\n\n")
	}
}
