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
		return nil, fmt.Errorf("failed to get document list: %w", err)
	}

	return documents, nil
}

func (repo *Repository) EditDocumentInfo(
	ctx context.Context,
	passengerID uuid.UUID,
	edited entity.Document) (*entity.Document, error) {
	ex := repo.CheckTx(ctx)

	var changedDocument entity.Document

	err := ex.GetContext(ctx, &changedDocument, EditDocumentInfoQuery,
		passengerID,
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

	return &changedDocument, nil
}

func (repo *Repository) RemoveDocumentInfo(ctx context.Context, documentID uuid.UUID) error {
	ex := repo.CheckTx(ctx)

	_, err := ex.QueryContext(
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
