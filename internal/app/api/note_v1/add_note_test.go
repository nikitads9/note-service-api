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

func Test_AddNote(t *testing.T) {
	var (
		ctx         = context.Background()
		mock        = gomock.NewController(t)
		noteId      = gofakeit.Int64()
		noteTitle   = gofakeit.BeerName()
		noteContent = gofakeit.BeerStyle()

		validNoteInfo = &model.NoteInfo{
			Title: sql.NullString{
				String: noteTitle,
				Valid:  true,
			},
			Content: sql.NullString{
				String: noteContent,
				Valid:  true,
			},
		}
		validReq = &desc.AddNoteRequest{
			Note: &desc.Notes{
				Title:   noteTitle,
				Content: noteContent,
			},
		}
	)
	noteRepoMock := noteRepoMocks.NewMockINoteRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().AddNote(ctx, validNoteInfo).Return(noteId, nil).Times(1),
		noteRepoMock.EXPECT().AddNote(ctx, validNoteInfo).Return(noteId, errors.New("some error")).Times(1),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.AddNote(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, res.GetId(), noteId)
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.AddNote(ctx, validReq)
		require.Error(t, err)
	})
}
