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
	"google.golang.org/protobuf/types/known/emptypb"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_UpdateNote(t *testing.T) {
	var (
		ctx          = context.Background()
		mock         = gomock.NewController(t)
		noteId       = gofakeit.Int64()
		noteTitle    = gofakeit.BeerName()
		noteContent  = gofakeit.BeerStyle()
		validRequest = &desc.UpdateNoteRequest{
			Id: noteId,
			Title: &wrapper.StringValue{
				Value: noteTitle,
			},
			Content: &wrapper.StringValue{
				Value: noteContent,
			},
		}
		validNote = &model.UpdateNoteInfo{
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
	)
	noteRepoMock := noteRepoMocks.NewMockINoteRepository(mock)
	gomock.InOrder(
		noteRepoMock.EXPECT().UpdateNote(ctx, validNote).Return(nil).Times(1),
		noteRepoMock.EXPECT().UpdateNote(ctx, validNote).Return(errors.New("someError")).Times(1),
	)

	api := newMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.UpdateNote(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.UpdateNote(ctx, validRequest)
		require.Error(t, err)
	})
}
