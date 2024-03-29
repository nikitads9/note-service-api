package note_v1

import (
	"context"
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/nikitads9/note-service-api/internal/model"
	noteRepoMocks "github.com/nikitads9/note-service-api/internal/repository/mocks"
	"github.com/nikitads9/note-service-api/internal/service/note"
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
		noteErr     = errors.New(gofakeit.Phrase())

		validNoteInfo = &model.NoteInfo{
			Title:   noteTitle,
			Content: noteContent,
		}
		validReq = &desc.AddNoteRequest{
			Note: &desc.Note{
				Title:   noteTitle,
				Content: noteContent,
			},
		}
	)
	noteRepoMock := noteRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().AddNote(ctx, validNoteInfo).Return(noteId, nil).Times(1),
		noteRepoMock.EXPECT().AddNote(ctx, validNoteInfo).Return(int64(0), noteErr).Times(1),
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
		require.Equal(t, err, noteErr)
	})
}
