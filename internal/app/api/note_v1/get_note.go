package note_v1

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/convert"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	note, err := i.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		NoteInfo: convert.ToDescNoteInfo(note),
	}, nil
}
