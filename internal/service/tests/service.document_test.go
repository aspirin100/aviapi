package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aspirin100/aviapi/internal/entity"
	"github.com/aspirin100/aviapi/internal/repository"
	"github.com/aspirin100/aviapi/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	docUids = []uuid.UUID{
		uuid.MustParse("4d6d45a8-a5cf-408a-908c-0d5d30d1e949"),
		uuid.MustParse("9279ae0d-e1cc-4389-9605-be31e8fcf3ac"),
		uuid.MustParse("843ca979-86a2-49ce-ba13-465840075976"),
	}
)

func initDocumentService() (*service.DocumentService, error) {
	repo, err := repository.NewConnection("postgres", PostgresDSN)
	if err != nil {
		return nil, err
	}

	return service.NewDocumentService(repo), nil
}

func TestEditDocumentInfo(t *testing.T) {
	srv, err := initDocumentService()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	type Params struct {
		id     uuid.UUID
		edited entity.Document
	}

	cases := []struct {
		Name        string
		ExpectedErr error
		Args        Params
	}{
		{
			Name:        "ok case",
			ExpectedErr: nil,
			Args: Params{
				id: docUids[0],
				edited: entity.Document{
					Type: "type from test",
				},
			},
		},
		{
			Name:        "not found case",
			ExpectedErr: entity.ErrDocumentNotFound,
			Args: Params{
				id: uuid.Nil,
				edited: entity.Document{
					Type: "type from test",
				},
			},
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.Name, func(t *testing.T) {
			_, err = srv.EditDocumentInfo(
				context.TODO(),
				tcase.Args.id,
				tcase.Args.edited)

			require.EqualValues(t, tcase.ExpectedErr, err)
		})
	}
}
