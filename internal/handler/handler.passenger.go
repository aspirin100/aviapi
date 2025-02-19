package handler

import (
	"fmt"
	"net/http"

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

}

func (h *Handler) RemovePassengerInfo(ctx *gin.Context) {

}
