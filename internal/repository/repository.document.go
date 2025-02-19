package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
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
		return nil, fmt.Errorf("failed to get document list: %w", err)
	}

	return documents, nil
}

func (repo *Repository) EditDocumentInfo(
	ctx context.Context,
	documentID uuid.UUID,
	edited entity.Document) (*entity.Document, error) {
	ex := repo.CheckTx(ctx)

	var changedDocument entity.Document

	err := ex.GetContext(ctx, &changedDocument, EditDocumentInfoQuery,
		documentID,
		edited.Type,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, entity.ErrDocumentNotFound
		default:
			return nil, fmt.Errorf("failed to edit document: %w", err)
		}
	}

	return &changedDocument, nil
}

func (repo *Repository) RemoveDocumentInfo(ctx context.Context, documentID uuid.UUID) error {
	ex := repo.CheckTx(ctx)

	_, err := ex.QueryContext( //nolint:gocritic
		ctx,
		RemoveDocumentInfoQuery,
		documentID) 
	if err != nil {
		return fmt.Errorf("failed to remove document info: %w", err)
	}

	return nil
}

const (
	GetDocumentListQuery = `
	SELECT
		document_type,
		id
	FROM documents
	WHERE
		passenger_id = $1; 
	`

	EditDocumentInfoQuery = `
		UPDATE
			documents
		SET
			document_type = CASE WHEN $2 = '' THEN document_type ELSE $2 END
		WHERE
			id = $1
		RETURNING
			id,
			document_type;
	`

	RemoveDocumentInfoQuery = `
		DELETE FROM documents
		WHERE
			id = $1;
	`
)
