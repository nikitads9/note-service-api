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

func Test_MultiAdd(t *testing.T) {
	var (
		ctx          = context.Background()
		mock         = gomock.NewController(t)
		noteTitle1   = gofakeit.BeerName()
		noteContent1 = gofakeit.BeerStyle()
		noteTitle2   = gofakeit.BeerName()
		noteContent2 = gofakeit.BeerStyle()

		validNotes = []*model.NoteInfo{
			{
				Title: sql.NullString{
					String: noteTitle1,
					Valid:  true,
				},
				Content: sql.NullString{
					String: noteContent1,
					Valid:  true,
				},
			},
			{
				Title: sql.NullString{
					String: noteTitle2,
					Valid:  true,
				},
				Content: sql.NullString{
					String: noteContent2,
					Valid:  true,
				},
			},
		}
		validReq = &desc.MultiAddRequest{
			Notes: []*desc.Notes{
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
	noteRepoMock := noteRepoMocks.NewMockINoteRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().MultiAdd(ctx, validNotes).Return(int64(len(validReq.GetNotes())), nil).Times(1),
		noteRepoMock.EXPECT().MultiAdd(ctx, validNotes).Return(int64(len(validReq.GetNotes())), errors.New("some error")).Times(1),
	)

	api := NewMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		added, err := api.MultiAdd(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, added.GetResult().GetCount(), int64(len(validReq.GetNotes())))
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.MultiAdd(ctx, validReq)
		require.Error(t, err)
	})
}
