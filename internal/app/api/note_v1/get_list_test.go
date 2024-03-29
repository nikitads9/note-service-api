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

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_GetList(t *testing.T) {
	var (
		ctx         = context.Background()
		mock        = gomock.NewController(t)
		noteId      = gofakeit.Int64()
		noteTitle   = gofakeit.BeerName()
		noteContent = gofakeit.BeerStyle()
		noteErr     = errors.New(gofakeit.Phrase())

		validResponse = []*model.NoteInfo{{
			Id:      noteId,
			Title:   noteTitle,
			Content: noteContent,
		},
			{
				Id:      noteId,
				Title:   noteTitle,
				Content: noteContent,
			},
		}
	)
	noteRepoMock := noteRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().GetList(ctx).Return(validResponse, nil).Times(1),
		noteRepoMock.EXPECT().GetList(ctx).Return(nil, noteErr).Times(1),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetList(ctx, &emptypb.Empty{})
		require.Nil(t, err)
		require.Condition(t, func() bool {
			return len(res.GetNoteInfo()) == len(validResponse)
		}, "the number of returned entries should be equal to length of valid response slice")
	})

	t.Run("error case", func(t *testing.T) {
		res, err := api.GetList(ctx, &emptypb.Empty{})
		require.Error(t, err)
		require.Equal(t, err, noteErr)
		require.Nil(t, res)
	})
}
