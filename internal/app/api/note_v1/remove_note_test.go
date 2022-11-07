package note_v1

import (
	"context"
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	noteRepoMocks "github.com/nikitads9/note-service-api/internal/repository/mocks"
	"github.com/nikitads9/note-service-api/internal/service/note"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	"github.com/stretchr/testify/require"
)

func Test_RemoveNote(t *testing.T) {
	var (
		ctx          = context.Background()
		mock         = gomock.NewController(t)
		noteId       = gofakeit.Int64()
		validRequest = &desc.RemoveNoteRequest{
			Id: noteId,
		}
	)
	noteRepoMock := noteRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().RemoveNote(ctx, noteId).Return(noteId, nil).Times(1),
		noteRepoMock.EXPECT().RemoveNote(ctx, noteId).Return(int64(0), errors.New("someError")).Times(1),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		resp, err := api.RemoveNote(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, validRequest.GetId(), resp.GetRemoved())
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.RemoveNote(ctx, validRequest)
		require.Error(t, err)
	})
}
