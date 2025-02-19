package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aspirin100/aviapi/internal/repository"
	"github.com/google/uuid"
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

func TestGetPassengerReport(t *testing.T) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	startDate, err := time.Parse(time.DateTime, "2023-11-10 12:00:00")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	endDate, err := time.Parse(time.DateTime, "2023-11-10 12:00:00")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	passengerID := uuid.MustParse("a3bb4716-f346-498e-8e89-b9db8474da24")

	report, err := repo.GetReport(
		context.Background(),
		passengerID,
		startDate,
		endDate)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, item := range report {
		fmt.Printf("Order ID: %s, From: %s, To: %s, Service Provided: %t\n",
			item.OrderID, item.FromCountry, item.ToCountry, item.ServiceProvided)
	}
}
