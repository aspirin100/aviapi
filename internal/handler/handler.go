package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aspirin100/aviapi/internal/config"
)

type Handler struct {
	airflightManager AirflightManager
	server           *http.Server
}

func New(airflightManager AirflightManager, cfg *config.Config) *Handler {
	return &Handler{
		airflightManager: airflightManager,
		server: &http.Server{
			Addr:         cfg.Hostname + ":" + cfg.Port,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (h *Handler) Start() error {
	err := h.server.ListenAndServe()
	if err != http.ErrServerClosed {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return nil
}

func (h *Handler) Shutdown(ctx context.Context) error {
	err := h.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop http server: %w", err)
	}

	return nil
}
