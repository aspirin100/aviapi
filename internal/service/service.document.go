package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type DocumentHandler interface {
	GetDocumentList(ctx context.Context, passengerID uuid.UUID) ([]entity.Document, error)
	EditDocumentInfo(ctx context.Context, documentID uuid.UUID, edited entity.Document) (*entity.Document, error)
	RemoveDocumentInfo(ctx context.Context, documentID uuid.UUID) error
	BeginTx(ctx context.Context) (context.Context, entity.CommitOrRollback, error)
}

type DocumentService struct {
	documentHandler DocumentHandler
}

func NewDocumentService(documentHandler DocumentHandler) *DocumentService {
	return &DocumentService{
		documentHandler: documentHandler,
	}
}

func (ds *DocumentService) GetDocumentList(
	ctx context.Context,
	passengerID uuid.UUID) ([]entity.Document, error) {
	const op = "service.GetDocumentList"

	ctx, commitOrRollback, err := ds.documentHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	documents, err := ds.documentHandler.GetDocumentList(ctx, passengerID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return documents, nil
}

func (ds *DocumentService) EditDocumentInfo(
	ctx context.Context,
	documentID uuid.UUID,
	edited entity.Document) (*entity.Document, error) {
	ctx, commitOrRollback, err := ds.documentHandler.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	changedDocument, err := ds.documentHandler.EditDocumentInfo(
		ctx,
		documentID,
		edited)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrDocumentNotFound):
			return nil, entity.ErrDocumentNotFound
		default:
			return nil, fmt.Errorf("failed to edit document: %w", err)
		}
	}

	return changedDocument, nil
}

func (ds *DocumentService) RemoveDocumentInfo(
	ctx context.Context,
	documentID uuid.UUID) error {
	const op = "service.RemoveDocumentInfo"

	ctx, commitOrRollback, err := ds.documentHandler.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(err error) {
		errTx := commitOrRollback(err)
		if errTx != nil {
			fmt.Printf("commit/rollback error: %v", errTx)
		}
	}(err)

	err = ds.documentHandler.RemoveDocumentInfo(
		ctx,
		documentID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
