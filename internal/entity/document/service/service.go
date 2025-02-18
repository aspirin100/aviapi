package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type DocumentHandler interface {
	GetDocumentList(ctx context.Context, passengerID uuid.UUID) ([]entity.Document, error)
	EditDocumentInfo(ctx context.Context, documentID uuid.UUID, edited entity.Document) (*entity.Document, error)
	RemoveDocumentInfo(documentID uuid.UUID) error
}

type DocumentService struct {
	documentHandler DocumentHandler
}

func New(documentHandler DocumentHandler) *DocumentService {
	return &DocumentService{
		documentHandler: documentHandler,
	}
}

func (ds *DocumentService) GetDocumentList(ctx context.Context, passengerID uuid.UUID) ([]entity.Document, error) {
	return nil, nil
}

func (ds *DocumentService) EditDocumentInfo(ctx context.Context, documentID uuid.UUID, edited entity.Document) (*entity.Document, error) {
	return nil, nil
}

func (ds *DocumentService) RemoveDocumentInfo(documentID uuid.UUID) error {
	return nil
}
