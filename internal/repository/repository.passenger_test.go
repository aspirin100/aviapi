package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetPassengerList(t *testing.T) {
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
			OrderID:     uuid.Nil,
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
