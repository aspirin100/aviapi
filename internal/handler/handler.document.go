package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetDocumentList(ctx *gin.Context) {
	parsedID, err := uuid.Parse(ctx.Param("passenger_id"))
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	documents, err := h.airflightManager.GetDocumentList(
		ctx,
		parsedID)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, documents)
}

func (h *Handler) EditDocumentInfo(ctx *gin.Context) {

}

func (h *Handler) RemoveDocumentInfo(ctx *gin.Context) {

}

// func validateEditDocumentRequest(ctx *gin.Context){

// }
