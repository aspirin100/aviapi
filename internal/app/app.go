package app

import (
	"fmt"

	"github.com/aspirin100/aviapi/internal/config"
	ticketservice "github.com/aspirin100/aviapi/internal/entity/airticket/service"
	docservice "github.com/aspirin100/aviapi/internal/entity/document/service"
	passervice "github.com/aspirin100/aviapi/internal/entity/passenger/service"
	"github.com/aspirin100/aviapi/internal/handler"
	"github.com/aspirin100/aviapi/internal/repository"
	infoservice "github.com/aspirin100/aviapi/internal/service"
)

const (
	databaseDriverName = "postgres"
)

type manager struct {
	*docservice.DocumentService
	*ticketservice.AirticketService
	*passervice.PassengerService
	*infoservice.Service
}

type App struct {
	serverHandler *handler.Handler
}

func New(cfg config.Config) (*App, error) {
	// repository constructor
	repo, err := repository.NewConnection(databaseDriverName, cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to create repo instance: %w", err)
	}

	// service layer constructor
	mng := initManager(repo)

	serverHandler := handler.New(mng, cfg)

	return &App{
		serverHandler: serverHandler,
	}, nil
}

func initManager(repo *repository.Repository) *manager {
	return &manager{
		docservice.New(repo),
		ticketservice.New(repo),
		passervice.New(repo),
		infoservice.New(repo),
	}
}
