package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetPassengerList(ctx *gin.Context) {
	parsedID, err := uuid.Parse(ctx.Param("order_id"))
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	passengers, err := h.airflightManager.GetPassengerList(
		ctx,
		parsedID)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, passengers)
}

func (h *Handler) EditPassengerInfo(ctx *gin.Context) {
	parsedID, editedPassenger, err := validateEditPassengerRequest(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	changedPassenger, err := h.airflightManager.EditPassengerInfo(
		ctx,
		*parsedID,
		*editedPassenger)
	if err != nil {
		fmt.Println(err)

		switch {
		case errors.Is(err, entity.ErrPassengerNotFound):
			ctx.Status(http.StatusNotFound)
		default:
			ctx.Status(http.StatusInternalServerError)
		}

		return
	}

	ctx.JSON(http.StatusOK, changedPassenger)
}

func (h *Handler) RemovePassengerInfo(ctx *gin.Context) {
	parsedID, err := uuid.Parse(ctx.Param("passenger_id"))
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	err = h.airflightManager.RemovePassengerInfo(ctx, parsedID)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func validateEditPassengerRequest(ctx *gin.Context) (*uuid.UUID, *entity.Passenger, error) {
	passenger_id := ctx.Param("passenger_id")

	parsedID, err := uuid.Parse(passenger_id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse passenger id: %w", err)
	}

	var editedPassenger entity.Passenger

	dec := json.NewDecoder(ctx.Request.Body)

	err = dec.Decode(&editedPassenger)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	return &parsedID, &editedPassenger, nil
}
