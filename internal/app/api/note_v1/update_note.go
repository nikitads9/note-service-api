package note_v1

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/convert"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	i.noteService.UpdateNote(ctx, convert.ToNoteUpdateInfo(req))

	return &emptypb.Empty{}, nil
}
