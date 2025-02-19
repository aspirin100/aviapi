package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/aspirin100/aviapi/internal/config"
	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	airflightManager AirflightManager
	server           *http.Server
}

func New(airflightManager AirflightManager, cfg *config.Config) *Handler {
	serverHandler := &Handler{
		airflightManager: airflightManager,
	}

	router := gin.Default()

	_ = router.GET("/airticket", serverHandler.GetTicketList)
	_ = router.PATCH("/airticket/:order_id", serverHandler.EditTicketInfo)
	_ = router.DELETE("/airticket/:order_id", serverHandler.RemoveTicketInfo)

	_ = router.GET("/:passenger_id/documents", serverHandler.GetDocumentList)
	_ = router.PATCH("/documents/:document_id", serverHandler.EditDocumentInfo)
	_ = router.DELETE("/documents/:document_id", serverHandler.RemoveDocumentInfo)

	_ = router.GET("/airticket/:order_id/passengers", serverHandler.GetPassengerList)
	_ = router.PATCH("/passengers/:passenger_id", serverHandler.EditPassengerInfo)
	_ = router.DELETE("/passengers/:passenger_id", serverHandler.RemovePassengerInfo)

	_ = router.GET("airticket/:order_id/info", serverHandler.GetFullInfo)

	serverHandler.server = &http.Server{
		Addr:         cfg.Hostname + ":" + cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		Handler:      router,
	}

	return serverHandler
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

func (h *Handler) GetFullInfo(ctx *gin.Context) {
	order_id := ctx.Param("order_id")

	parsedID, err := uuid.Parse(order_id)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	fullinfo, err := h.airflightManager.GetFullInfo(ctx, parsedID)
	if err != nil {
		fmt.Println(err)

		switch {
		case errors.Is(err, entity.ErrTicketNotFound):
			ctx.Status(http.StatusNotFound)
		default:
			ctx.Status(http.StatusInternalServerError)
		}

		return
	}

	ctx.JSON(http.StatusOK, fullinfo)
}

