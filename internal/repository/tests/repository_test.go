package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aspirin100/aviapi/internal/repository"
	"github.com/google/uuid"
)

func TestGetFullInfo(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	infoList, err := repo.GetFullInfo(context.Background(), uuid.Nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, val := range infoList {
		fmt.Print(val, "\n\n")
	}
}
