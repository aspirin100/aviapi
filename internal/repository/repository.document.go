package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

var (
	ErrDocumentNotFound = errors.New("document was not found")
)

func (repo *Repository) GetDocumentList(ctx context.Context, passengerID uuid.UUID) ([]entity.Document, error) {
	ex := repo.CheckTx(ctx)

	documents := []entity.Document{}

	err := ex.SelectContext(
		ctx,
		&documents,
		GetDocumentListQuery,
		passengerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tickets list: %w", err)
	}

	return documents, nil
}

func (repo *Repository) EditDocumentInfo(
	ctx context.Context,
	documentID uuid.UUID,
	edited entity.Document) (*entity.Document, error) {
	ex := repo.CheckTx(ctx)

	var finalDocument entity.Document

	err := ex.GetContext(ctx, &finalDocument, EditDocumentInfoQuery,
		documentID,
		edited.Type,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrDocumentNotFound
		default:
			return nil, fmt.Errorf("failed to edit document: %w", err)
		}
	}

	return &finalDocument, nil
}

func (repo *Repository) RemoveDocumentInfo(ctx context.Context, documentID uuid.UUID) error {
	ex := repo.CheckTx(ctx)

	_, err := ex.QueryContext(
		ctx,
		RemoveDocumentInfoQuery,
		documentID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrDocumentNotFound
		default:
			return fmt.Errorf("failed to remove ticket info: %w", err)
		}
	}

	return nil
}

const (
	GetDocumentListQuery = `
	SELECT
		type,
		id
	FROM documents
	WHERE
		passenger_id = $1; 
	`

	EditDocumentInfoQuery = `
		UPDATE
			documents
		SET
			type = $2
		WHERE
			id = $1;
	`

	RemoveDocumentInfoQuery = `
		DELETE FROM documents
		WHERE
			id = $1;
	`
)
