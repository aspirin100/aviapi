package repository

import (
	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/google/uuid"
)

func (repo *Repository) GetDocumentList(passengerID uuid.UUID) ([]entity.Document, error) {
	return nil, nil
}

func (repo *Repository) EditDocumentInfo(documentID uuid.UUID, edited entity.Document) error {
	return nil
}

func (repo *Repository) RemoveDocumentInfo(documentID uuid.UUID) error {
	return nil
}
