package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/nikitads9/note-service-api/internal/app/model"
	noteRepoMocks "github.com/nikitads9/note-service-api/internal/app/repository/mocks"
	"github.com/nikitads9/note-service-api/internal/app/service/note"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	"github.com/stretchr/testify/require"
)

func Test_GetNote(t *testing.T) {
	var (
		ctx          = context.Background()
		mock         = gomock.NewController(t)
		noteId       = gofakeit.Int64()
		noteTitle    = gofakeit.BeerName()
		noteContent  = gofakeit.BeerStyle()
		validRequest = &desc.GetNoteRequest{
			Id: noteId,
		}
		validResponse = &model.NoteInfo{
			Id: noteId,
			Title: sql.NullString{
				String: noteTitle,
				Valid:  true,
			},
			Content: sql.NullString{
				String: noteContent,
				Valid:  true,
			},
		}
		errRepo = errors.New("ebanyi rot etogo kasino")
	)
	noteRepoMock := noteRepoMocks.NewMockINoteRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().GetNote(ctx, noteId).Return(validResponse, nil).Times(1),
		noteRepoMock.EXPECT().GetNote(ctx, noteId).Return(nil, errRepo).Times(1),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		resp, err := api.GetNote(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, resp.GetId(), validRequest.GetId())
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.GetNote(ctx, validRequest)
		require.Error(t, err)
		require.Equal(t, errRepo, err)
	})
}
