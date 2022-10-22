package note_v1

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nikitads9/note-service-api/internal/app/convert"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) AddNote(ctx context.Context, req *desc.AddNoteRequest) (*desc.AddNoteResponse, error) {
	id, err := i.noteService.AddNote(ctx, convert.ToNoteInfo(&desc.Notes{
		Title:   req.GetNote().GetTitle(),
		Content: req.GetNote().GetContent(),
	}))
	if err != nil {
		return nil, err
	}

	return &desc.AddNoteResponse{
		Id: id,
	}, nil
}
