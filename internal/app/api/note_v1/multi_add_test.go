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

func Test_MultiAdd(t *testing.T) {
	var (
		ctx          = context.Background()
		mock         = gomock.NewController(t)
		noteTitle1   = gofakeit.BeerName()
		noteContent1 = gofakeit.BeerStyle()
		noteTitle2   = gofakeit.BeerName()
		noteContent2 = gofakeit.BeerStyle()
		noteErr      = errors.New(gofakeit.Phrase())

		validNotes = []*model.NoteInfo{
			{
				Title: noteTitle1,

				Content: noteContent1,
			},
			{
				Title:   noteTitle2,
				Content: noteContent2,
			},
		}
		validReq = &desc.MultiAddRequest{
			Notes: []*desc.Note{
				{
					Title:   noteTitle1,
					Content: noteContent1,
				},
				{
					Title:   noteTitle2,
					Content: noteContent2,
				},
			},
		}
	)
	noteRepoMock := noteRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().MultiAdd(ctx, validNotes).Return(int64(len(validReq.GetNotes())), nil).Times(1),
		noteRepoMock.EXPECT().MultiAdd(ctx, validNotes).Return(int64(0), noteErr).Times(1),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		added, err := api.MultiAdd(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, added.GetCount(), int64(len(validReq.GetNotes())))
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.MultiAdd(ctx, validReq)
		require.Error(t, err)
		require.Equal(t, err, noteErr)
	})
}
