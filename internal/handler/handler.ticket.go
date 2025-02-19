package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetTicketList(ctx *gin.Context) {
	tickets, err := h.airflightManager.GetTicketList(ctx)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (h *Handler) EditTicketInfo(ctx *gin.Context) {
	orderID, editedTicket, err := validateEditTicketRequest(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	spew.Dump(editedTicket)

	changedTicket, err := h.airflightManager.EditTicketInfo(
		ctx,
		*orderID,
		*editedTicket)
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

	ctx.JSON(http.StatusOK, changedTicket)
}

func (h *Handler) RemoveTicketInfo(ctx *gin.Context) {
	parsedID, err := uuid.Parse(ctx.Param("order_id"))
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	err = h.airflightManager.RemoveTicketInfo(ctx, parsedID)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func validateEditTicketRequest(ctx *gin.Context) (*uuid.UUID, *entity.AirTicket, error) {
	order_id := ctx.Param("order_id")

	parsedID, err := uuid.Parse(order_id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse order id: %w", err)
	}

	var editedTicket entity.AirTicket

	dec := json.NewDecoder(ctx.Request.Body)

	err = dec.Decode(&editedTicket)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	return &parsedID, &editedTicket, nil
}
