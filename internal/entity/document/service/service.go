package service

import (
	"github.com/google/uuid"

	"github.com/aspirin100/aviapi/internal/entity"
)

type DocumentHandler interface {
	GetDocumentList(passengerID uuid.UUID) ([]entity.Document, error)
	EditDocumentInfo(documentID uuid.UUID, edited entity.Document) error
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

func (ds *DocumentService) GetDocumentList(passengerID uuid.UUID) ([]entity.Document, error) {
	return nil, nil
}

func (ds *DocumentService) EditDocumentInfo(documentID uuid.UUID, edited entity.Document) error {
	return nil
}

func (ds *DocumentService) RemoveDocumentInfo(documentID uuid.UUID) error {
	return nil
}
