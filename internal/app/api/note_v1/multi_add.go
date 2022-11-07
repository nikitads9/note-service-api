package note_v1

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/convert"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) MultiAdd(ctx context.Context, req *desc.MultiAddRequest) (*desc.MultiAddResponse, error) {
	res, err := i.noteService.MultiAdd(ctx, convert.ToNotesInfo(req.GetNotes()))
	if err != nil {
		return nil, err
	}

	return &desc.MultiAddResponse{
		Count: res,
	}, nil
}
