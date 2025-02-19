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
	parsedID, editedDocument, err := validateEditDocumentRequest(ctx)
	if err != nil {
		fmt.Println(err)

		ctx.Status(http.StatusBadRequest)

		return
	}

	changedDocument, err := h.airflightManager.EditDocumentInfo(
		ctx,
		*parsedID,
		*editedDocument)
	if err != nil {
		fmt.Println(err)

		switch {
		case errors.Is(err, entity.ErrDocumentNotFound):
			ctx.Status(http.StatusNotFound)
		default:
			ctx.Status(http.StatusInternalServerError)
		}

		return
	}

	ctx.JSON(http.StatusOK, changedDocument)
}

func (h *Handler) RemoveDocumentInfo(ctx *gin.Context) {

}

func validateEditDocumentRequest(ctx *gin.Context) (*uuid.UUID, *entity.Document, error) {
	document_id := ctx.Param("document_id")

	parsedID, err := uuid.Parse(document_id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse document id: %w", err)
	}

	var editedDocument entity.Document

	dec := json.NewDecoder(ctx.Request.Body)

	err = dec.Decode(&editedDocument)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	return &parsedID, &editedDocument, nil
}
