package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/config"
	"github.com/aspirin100/aviapi/internal/entity"
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
	_ = router.GET("/passengers/:passenger_id/report", serverHandler.GetReport)

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
	if !errors.Is(err, http.ErrServerClosed) {
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
	orderID := ctx.Param("order_id")

	parsedID, err := uuid.Parse(orderID)
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

type getReportArgs struct {
	passengerID uuid.UUID
	startPeriod time.Time
	endPeriod   time.Time
}

func (h *Handler) GetReport(ctx *gin.Context) {
	args, err := validateGetReportRequest(ctx)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	report, err := h.airflightManager.GetReport(
		ctx,
		args.passengerID,
		args.startPeriod,
		args.endPeriod)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, report)
}

func validateGetReportRequest(ctx *gin.Context) (*getReportArgs, error) {
	parsedID, err := uuid.Parse(ctx.Param("passenger_id"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse passenger id: %w", err)
	}

	period := struct {
		StartPeriod string `json:"start_period"`
		EndPeriod   string `json:"end_period"`
	}{}

	dec := json.NewDecoder(ctx.Request.Body)

	err = dec.Decode(&period)
	if err != nil {
		return nil, fmt.Errorf("get report request body decode fail: %w", err)
	}

	startPeriod, err := time.Parse(time.DateTime, period.StartPeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to parse time: %w", err)
	}

	endPeriod, err := time.Parse(time.DateTime, period.EndPeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to parse time: %w", err)
	}

	return &getReportArgs{
		passengerID: parsedID,
		startPeriod: startPeriod,
		endPeriod:   endPeriod,
	}, nil
}
