package note_v1

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/convert"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) MultiAdd(ctx context.Context, req *desc.MultiAddRequest) (*desc.MultiAddResponse, error) {
	res, err := i.noteService.MultiAdd(ctx, convert.ToNotesInfo(req))
	if err != nil {
		return nil, err
	}

	return &desc.MultiAddResponse{
		Result: &desc.MultiAddResponse_Result{
			Count: res,
		},
	}, nil
}
